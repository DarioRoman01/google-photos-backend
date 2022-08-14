package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DarioRoman01/photos/mailpb"
	"github.com/mailgun/mailgun-go/v3"
)

type Server struct {
	mailpb.UnimplementedMailServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) CreateMessage(r *mailpb.SendMailRequest, mg *mailgun.MailgunImpl) (*mailgun.Message, error) {
	msg := mg.NewMessage(
		fmt.Sprintf("mailgun@%s", os.Getenv("MAILGUN_DOMAIN")),
		r.GetSubject(),
		r.GetBody(),
		r.GetReceiver(),
	)

	msg.SetTemplate(r.GetType())
	msg.AddTemplateVariable("token", r.GetToken())
	msg.AddTemplateVariable("username", r.GetUser())
	return msg, nil
}

func (s *Server) SendMail(ctx context.Context, in *mailpb.SendMailRequest) (*mailpb.SendMailResponse, error) {
	mg := mailgun.NewMailgun(os.Getenv("MAILGUN_DOMAIN"), os.Getenv("MAILGUN_API_KEY"))
	msg, err := s.CreateMessage(in, mg)
	if err != nil {
		return nil, err
	}

	_, id, err := mg.Send(ctx, msg)
	if err != nil {
		return nil, err
	}

	return &mailpb.SendMailResponse{Id: id}, nil
}
