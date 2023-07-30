package ftp

import (
	"context"
	"github.com/jlaffaye/ftp"
	"github.com/spf13/viper"
)

type Ftp struct {
	Connection *ftp.ServerConn
	ctx        context.Context
}

func Init(ctx context.Context) (*Ftp, error) {
	c, err := ftp.Dial(viper.GetString("FTP_HOST"))

	if err != nil {
		return nil, err
	}

	err = c.Login(viper.GetString("FTP_USER"), viper.GetString("FTP_PASSWORD"))

	if err != nil {
		return nil, err
	}

	return &Ftp{c, ctx}, nil
}

func (f *Ftp) Quit() error {
	return f.Connection.Quit()
}
