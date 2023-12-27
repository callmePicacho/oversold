package app

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"oversold/util"
	"oversold/util/model"
	"oversold/util/mutex"
)

func OversoldMutex(c *gin.Context) {
	// 获取请求参数
	req := Marshal(c)

	mutex.OversoldMutex.Lock()
	defer mutex.OversoldMutex.Unlock()

	// 查询库存 -> 扣减库存 -> 产生订单
	err := util.GetMysqlConn().Transaction(func(tx *gorm.DB) error {
		// 查询库存
		var product model.Product
		// select * from product where sku_id = ?
		err := tx.Where("sku_id = ?", req.SkuId).First(&product).Error
		if err != nil {
			return err
		}

		if product.Num >= req.Num { // 库存足够
			// 扣减库存
			product.Num -= req.Num
			// update product set num =? where id = ?
			err = tx.Model(&model.Product{}).
				Where("id = ?", product.ID).
				Update("num", product.Num).Error
			if err != nil {
				return err
			}

			// 创建订单
			// insert into order(user_id, product_id) values(?,?)
			err = tx.Create(&model.Order{UserID: req.UserId, ProductID: req.SkuId}).Error
			if err != nil {
				return err
			}
		} else {
			return errors.New("库存不足")
		}

		return nil
	})

	if err != nil {
		Fail(c, fmt.Sprintf("下单失败，失败原因:%v", err.Error()))
		return
	}

	OK(c)
}
