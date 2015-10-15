package goluago

import (
	"github.com/mtabini/go-lua"
	"github.com/mtabini/goluago/pkg/crypto/hmac"
	"github.com/mtabini/goluago/pkg/encoding/base64"
	"github.com/mtabini/goluago/pkg/encoding/json"
	"github.com/mtabini/goluago/pkg/env"
	"github.com/mtabini/goluago/pkg/fmt"
	"github.com/mtabini/goluago/pkg/net/url"
	"github.com/mtabini/goluago/pkg/regexp"
	"github.com/mtabini/goluago/pkg/strings"
	"github.com/mtabini/goluago/pkg/time"
	"github.com/mtabini/goluago/pkg/uuid"
	"github.com/mtabini/goluago/util"
)

func Open(l *lua.State) {
	regexp.Open(l)
	strings.Open(l)
	json.Open(l)
	time.Open(l)
	fmt.Open(l)
	url.Open(l)
	util.Open(l)
	hmac.Open(l)
	base64.Open(l)
	env.Open(l)
	uuid.Open(l)
}
