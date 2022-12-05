package dao

import (
	"fmt"
	"ginapis/pkg/config"
	"ginapis/pkg/model"
	"github.com/bbdshow/bkit/caches"
	"github.com/bbdshow/bkit/db/mysql"
	"xorm.io/xorm"
)

type Dao struct {
	cfg      *config.Config
	mysql    *mysql.Xorm
	memCache caches.Cacher
}

func New(cfg *config.Config) *Dao {
	d := &Dao{
		cfg:      cfg,
		mysql:    mysql.NewXorm(cfg.Mysql),
		memCache: caches.NewLRUMemory(10000),
	}

	// 同步表和索引
	d.Sync2()

	return d
}

func (d *Dao) Close() {
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	if d.mysql != nil {
		_ = d.mysql.Close()
	}
}

func (d *Dao) Sync2() {
	if !d.cfg.Release() {
		err := d.mysql.Sync2()
		if err != nil {
			panic(err)
		}
		if err := d.mysql.Sync2(model.Models...); err != nil {
			panic(err)
		}
	}
}

func (d *Dao) MysqlTransaction(tx func(sess *xorm.Session) error) error {
	return d.mysql.Transaction(tx)
}

func (d *Dao) TxInsertOne(sess *xorm.Session, bean interface{}) error {
	affected, err := sess.Insert(bean)
	if affected != 1 {
		return fmt.Errorf("wrong affected:%v", affected)
	}
	return err
}

func (d *Dao) TxInserts(sess *xorm.Session, beans interface{}) error {
	_, err := sess.Insert(beans)
	return err
}

func (d *Dao) Insert(sess *xorm.Session, bean interface{}) error {
	var (
		affected int64
		err      error
	)
	if sess != nil {
		affected, err = sess.Insert(bean)
	} else {
		affected, err = d.mysql.Insert(bean)
	}

	if err != nil {
		return err
	}
	if affected != 1 {
		return fmt.Errorf("wrong affected:%d", affected)
	}
	return nil
}

func (d *Dao) Inserts(sess *xorm.Session, beans interface{}) (int64, error) {
	var (
		affected int64
		err      error
	)
	if sess != nil {
		affected, err = sess.Insert(beans)
	} else {
		affected, err = d.mysql.Insert(beans)
	}

	if err != nil {
		return affected, err
	}
	return affected, nil
}

func (d *Dao) Delete(sess *xorm.Session, id int64, bean interface{}) error {
	var (
		affected int64
		err      error
	)
	if sess != nil {
		affected, err = sess.ID(id).Delete(bean)
	} else {
		affected, err = d.mysql.ID(id).Delete(bean)
	}

	if err != nil {
		return err
	}
	if affected != 1 {
		return fmt.Errorf("wrong affected:%d", affected)
	}
	return nil
}
