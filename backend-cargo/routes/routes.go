package routes

import (
	"fmt"
	"net/http"

	"backend/dtos"
	"backend/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Routes(router *gin.Engine) {
	router.GET("/", welcome)
	router.POST("/v1/create-account", CreateAccount)
	router.POST("/v1/login", Login)
	router.GET("/v1/get-customer", GetCustomers)
	router.POST("/v1/account/:aid/booking", CreateBookingRequest)//(req quote)
	router.PUT("/v1/account/:aid/booking/:bid", UpdateBookingRequest)
	router.GET("/v1/account/:aid/booking/:bid", GetBookingRequest)//(single qiotelist)
	router.GET("/v1/account/:aid/bookings", GetAllBookingRequest)//(get quotelist)
	router.POST("/v1/account/:aid/booking/:bid/quote", CreateQuote)//(buy sell)
	router.PUT("/v1/account/:aid/booking/:bid/quote/:qid", UpdateQuote)
	router.GET("/v1/account/:aid/booking/:bid/quote/:qid", GetBookingQuote)
	router.GET("/v1/account/:aid/booking/:bid/quote", GetBookingAllQuotes)//(multi quote)
	router.PUT("/v1/account/:aid/booking/:bid/task/:tid", UpdateBookingTask)
	router.GET("/v1/account/:aid/booking/:bid/task/:tid", GetBookingTask)
	router.GET("/v1/account/:aid/booking/:bid/tasks", GetBookingAllTask)

	router.NoRoute(notFound)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	})
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
}

func CreateAccount(c *gin.Context) {
	req := &dtos.Account{}
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	val, err := service.CreateAccount(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": &val,
	})
}

func Login(c *gin.Context) {
	req := &dtos.Login{}
	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	res, err := service.Login(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": res,
	})
}

func GetCustomers(c *gin.Context) {
	res, err := service.GetCustomers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": res,
	})
}

func CreateBookingRequest(c *gin.Context) {
	// aid, err := uuid.Parse(c.Param("aid"))
	// if err != nil {
	// 	c.JSON(http.StatusBadGateway, gin.H{
	// 		"message": err,
	// 	})
	// }
	req := &dtos.Booking{}
	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	id, err := service.CreateBookingRequest(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": &id,
	})
}

func UpdateBookingRequest(c *gin.Context) {
	aid, err := uuid.Parse(c.Param("aid"))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err,
		})
		return
	}
	bid, err := uuid.Parse(c.Param("bid"))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err,
		})
		return
	}
	req := &dtos.Booking{}
	err = c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	req.Id = bid
	by := aid
	err = service.UpdateBookingRequest(req, by)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func GetBookingRequest(c *gin.Context) {
	// aid, err := uuid.Parse(c.Param("aid"))
	// if err != nil {
	// 	c.JSON(http.StatusBadGateway, gin.H{
	// 		"message": err,
	// 	})
	// }
	bid, err := uuid.Parse(c.Param("bid"))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err,
		})
		return
	}
	res, err := service.GetBookingRequest(bid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": res,
	})
}

func GetAllBookingRequest(c *gin.Context) {
	// aid, err := uuid.Parse(c.Param("aid"))
	// if err != nil {
	// 	c.JSON(http.StatusBadGateway, gin.H{
	// 		"message": err,
	// 	})
	// }
	res, err := service.AllBookingRequest()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": res,
	})
}

func CreateQuote(c *gin.Context) {

	bid, err := uuid.Parse(c.Param("bid"))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "invalid uuid",
		})
		return
	}
	req := &dtos.Quote{}
	err = c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "binding error",
		})
		return
	}
	req.BookingId = bid
	id, err := service.CreateQuote(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": &id,
	})
}

func UpdateQuote(c *gin.Context) {
	// aid, err := uuid.Parse(c.Param("aid"))
	// if err != nil {
	// 	c.JSON(http.StatusBadGateway, gin.H{
	// 		"message": err,
	// 	})
	// }
	bid, err := uuid.Parse(c.Param("bid"))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err,
		})
		return
	}
	qid, err := uuid.Parse(c.Param("qid"))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err,
		})
		return
	}
	req := &dtos.Quote{}
	err = c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	req.BookingId = bid
	req.ID = qid
	err = service.UpdateQuote(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func GetBookingQuote(c *gin.Context) {
	// aid, err := uuid.Parse(c.Param("aid"))
	// if err != nil {
	// 	c.JSON(http.StatusBadGateway, gin.H{
	// 		"message": err,
	// 	})
	// }
	// bid, err := uuid.Parse(c.Param("bid"))
	// if err != nil {
	// 	c.JSON(http.StatusBadGateway, gin.H{
	// 		"message": err,
	// 	})
	// }
	qid, err := uuid.Parse(c.Param("qid"))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err,
		})
		return
	}
	res, err := service.GetBookingQuote(qid)
	fmt.Println(res)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"err":     "db error",
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": res,
	})
}

func GetBookingAllQuotes(c *gin.Context) {
	// aid, err := uuid.Parse(c.Param("aid"))
	// if err != nil {
	// 	c.JSON(http.StatusBadGateway, gin.H{
	// 		"message": err,
	// 	})
	// }
	bid, err := uuid.Parse(c.Param("bid"))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err,
		})
		return
	}
	// qid, err := uuid.Parse(c.Param("qid"))
	// if err != nil {
	// 	c.JSON(http.StatusBadGateway, gin.H{
	// 		"message": err,
	// 	})
	// }
	res, err := service.GetBookingAllQuotes(bid)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"err":     "db error",
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": res,
	})
}

func GetBookingTask(c *gin.Context) {
	// aid, err := uuid.Parse(c.Param("aid"))
	// if err != nil {
	// 	c.JSON(http.StatusBadGateway, gin.H{
	// 		"message": err,
	// 	})
	// }
	// bid, err := uuid.Parse(c.Param("bid"))
	// if err != nil {
	// 	c.JSON(http.StatusBadGateway, gin.H{
	// 		"message": err,
	// 	})
	// }
	tid, err := uuid.Parse(c.Param("tid"))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err,
		})
		return
	}
	res, err := service.GetBookingTask(tid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": res,
	})
}

func GetBookingAllTask(c *gin.Context) {
	// aid, err := uuid.Parse(c.Param("aid"))
	// if err != nil {
	// 	c.JSON(http.StatusBadGateway, gin.H{
	// 		"message": err,
	// 	})
	// }
	bid, err := uuid.Parse(c.Param("bid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "invaild bid",
		})
		return
	}
	res, err := service.GetBookingAllTask(bid)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"err":     "db error",
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": res,
	})
}

func UpdateBookingTask(c *gin.Context) {
	aid, err := uuid.Parse(c.Param("aid"))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err,
		})
		return
	}
	bid, err := uuid.Parse(c.Param("bid"))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err,
		})
		return
	}
	tid, err := uuid.Parse(c.Param("tid"))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err,
		})
		return
	}

	req := &dtos.Task{}
	err = c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	req.BookingId = bid
	req.Id = tid
	by := aid
	err = service.UpdateBookingTask(req, by)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
