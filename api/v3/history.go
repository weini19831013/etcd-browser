package v3

import (
	"github.com/gin-gonic/gin"
	"github.com/wudaoluo/etcd-browser/etcdlib"
	"github.com/wudaoluo/etcd-browser/model"
	"net/http"
)


func History(c *gin.Context) {
	key := c.Param("action")
	etcdKey, _, err := etcdlib.EnsureKey(key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}

	node, err := etcdlib.Get(key)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"err": err.Error()})
		return
	}


	record := model.Get(etcdKey,node.IsDir)
	c.JSON(http.StatusOK, gin.H{"action": "history", "node": record,"key":key})
}