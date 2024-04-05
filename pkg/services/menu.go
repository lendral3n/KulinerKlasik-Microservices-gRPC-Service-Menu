package services

import (
	"context"
	"lendral3n/KulinerKlasik-Microservices-gRPC-Menu/pkg/db"
	"lendral3n/KulinerKlasik-Microservices-gRPC-Menu/pkg/models"
	"lendral3n/KulinerKlasik-Microservices-gRPC-Menu/pkg/pb"
	"net/http"
)

type Server struct {
	H db.Handler
	pb.UnimplementedMenuServiceServer
}

func (s *Server) CreateMenu(ctx context.Context, req *pb.CreateMenuRequest) (*pb.CreateMenuResponse, error) {
	var menu models.Menu

	menu.Name = req.Name
	menu.Price = req.Price
	menu.Stock = req.Stock

	if result := s.H.DB.Create(&menu); result.Error != nil {
		return &pb.CreateMenuResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.CreateMenuResponse{
		Status: http.StatusCreated,
		Id:     menu.Id,
	}, nil
}

func (s *Server) FineOne(ctx context.Context, req *pb.FindOneRequest) (*pb.FindOneResponse, error) {
	var menu models.Menu

	if result := s.H.DB.First(&menu, req.Id); result.Error != nil {
		return &pb.FindOneResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	data := &pb.FindOneData{
		Id:    menu.Id,
		Name:  menu.Name,
		Price: menu.Price,
		Stock: menu.Stock,
	}

	return &pb.FindOneResponse{
		Status: http.StatusOK,
		Data:   data,
	}, nil
}

func (s *Server) DecreaseStock(ctx context.Context, req *pb.DecreaseStockRequest) (*pb.DecreaseStockResponse, error) {
	var menu models.Menu

	if result := s.H.DB.First(&menu, req.Id); result.Error != nil {
		return &pb.DecreaseStockResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	if menu.Stock <= 0 {
		return &pb.DecreaseStockResponse{
			Status: http.StatusConflict,
			Error:  "Stock too low",
		}, nil
	}

	var log models.StockDecreaseLog

	if result := s.H.DB.Where(&models.StockDecreaseLog{
		OrderId: req.OrderId,
	}).First(&log); result.Error == nil {
		return &pb.DecreaseStockResponse{
			Status: http.StatusConflict,
			Error:  "Stock Already decreased",
		}, nil
	}

	menu.Stock = menu.Stock - 1
	s.H.DB.Save(&menu)

	log.OrderId = req.OrderId
	log.MenuRefer = menu.Id
	s.H.DB.Create(&log)

	return &pb.DecreaseStockResponse{
		Status: http.StatusOK,
	}, nil
}
