// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"ywadmin-v3/common/xerr"
	"ywadmin-v3/common/xfilters"
	"ywadmin-v3/service/admin/rpc/adminclient"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	sysStgroupUserFieldNames          = builder.RawFieldNames(&SysStgroupUser{})
	sysStgroupUserRows                = strings.Join(sysStgroupUserFieldNames, ",")
	sysStgroupUserRowsExpectAutoSet   = strings.Join(stringx.Remove(sysStgroupUserFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	sysStgroupUserRowsWithPlaceHolder = strings.Join(stringx.Remove(sysStgroupUserFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	sysStgroupUserModel interface {
		Insert(ctx context.Context, data *SysStgroupUser) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysStgroupUser, error)
		Update(ctx context.Context, data *SysStgroupUser) error
		Delete(ctx context.Context, id int64) error
		TransactInsert(ctx context.Context, data *adminclient.PolicyAssociatedUsersReq, opts ...string) error
		FindAll(ctx context.Context, filters ...interface{}) (*[]SysStgroupUser, error)
	}

	defaultSysStgroupUserModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysStgroupUser struct {
		Id        int64 `db:"id"`
		StgroupId int64 `db:"stgroup_id"` // 策略组id
		UserId    int64 `db:"user_id"`    // 用户id
	}
)

func newSysStgroupUserModel(conn sqlx.SqlConn) *defaultSysStgroupUserModel {
	return &defaultSysStgroupUserModel{
		conn:  conn,
		table: "`sys_stgroup_user`",
	}
}

func (m *defaultSysStgroupUserModel) Insert(ctx context.Context, data *SysStgroupUser) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, sysStgroupUserRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.StgroupId, data.UserId)
	return ret, err
}

func (m *defaultSysStgroupUserModel) FindOne(ctx context.Context, id int64) (*SysStgroupUser, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysStgroupUserRows, m.table)
	var resp SysStgroupUser
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysStgroupUserModel) FindAll(ctx context.Context, filters ...interface{}) (*[]SysStgroupUser, error) {
	filter := xfilters.Xfilters(filters...)
	var conditions string
	if len(filter) != 0 {
		conditions = "where " + filter
	}
	query := fmt.Sprintf("select %s from %s %s", sysStgroupUserRows, m.table, conditions)
	//fmt.Println(query)
	var resp []SysStgroupUser
	err := m.conn.QueryRowsCtx(ctx, &resp, query)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysStgroupUserModel) Update(ctx context.Context, data *SysStgroupUser) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysStgroupUserRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.StgroupId, data.UserId, data.Id)
	return err
}

func (m *defaultSysStgroupUserModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultSysStgroupUserModel) tableName() string {
	return m.table
}

func (m *defaultSysStgroupUserModel) TransactInsert(ctx context.Context, data *adminclient.PolicyAssociatedUsersReq, opts ...string) error {
	insertsql := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, sysStgroupUserRowsExpectAutoSet)
	if err := m.conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		stmt, err := session.Prepare(insertsql)
		if err != nil {
			return err
		}
		defer stmt.Close()
		tmp := fmt.Sprintf("delete from sys_stgroup_user where stgroup_id = %d", data.StgroupId)
		if len(opts) != 0 {
			tmp = fmt.Sprintf("delete from sys_stgroup_user where user_id = %d", data.StgroupId)
		}
		_, err = m.conn.ExecCtx(ctx, tmp)
		if err != nil {
			return xerr.NewErrMsg("删除用户与策略关联失败")
		}
		for _, v := range strings.Split(data.UserCheck, ",") {
			if v != "" {
				if len(opts) != 0 {
					_, err = stmt.ExecCtx(ctx, v, data.StgroupId)
				} else {
					_, err = stmt.ExecCtx(ctx, data.StgroupId, v)
				}
				if err != nil {
					return xerr.NewErrMsg("新增用户与策略关联失败")
				}
			}

		}
		return nil
	}); err != nil {
		return xerr.NewErrMsg(fmt.Sprintf("%+v", err.Error()))
	}
	return nil
}
