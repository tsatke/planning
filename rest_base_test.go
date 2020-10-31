package planning

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/suite"
	"github.com/tsatke/planning/db"
	"github.com/tsatke/planning/server"
)

func TestRestSuite(t *testing.T) {
	suite.Run(t, &RestSuite{})
}

type RestSuite struct {
	suite.Suite

	server *server.Server
}

func (suite *RestSuite) SetupTest() {
	testLog := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().
		Timestamp().
		Logger()

	db, err := db.Open(testLog, ":memory:")
	suite.NoError(err)

	suite.server = server.New(testLog, ":0", dataAccess{db})
	go func() {
		err := suite.server.Start(false)
		suite.NoError(err)
	}()
	<-suite.server.Listening() // wait for server startup to complete
}

func (suite *RestSuite) TearDownTest() {
	err := suite.server.Close()
	suite.NoError(err)
}

type (
	M map[string]interface{}

	TestRequest struct {
		suite    *RestSuite
		Method   string
		Endpoint string
		Data     M
	}
)

func (suite *RestSuite) Get(endpoint string) TestRequest {
	return suite.prepareRequest("GET", endpoint, nil)
}

func (suite *RestSuite) Post(endpoint string, data M) TestRequest {
	return suite.prepareRequest("POST", endpoint, data)
}

func (suite *RestSuite) prepareRequest(method, endpoint string, data M) TestRequest {
	return TestRequest{
		suite:    suite,
		Method:   method,
		Endpoint: endpoint,
		Data:     data,
	}
}

func (r TestRequest) Expect(rc int, data M) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(r.Data)
	r.suite.NoError(err)
	if err != nil {
		return // no point in continuing
	}

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	request, err := http.NewRequest(r.Method, "http://"+r.suite.server.Addr()+"/rest"+r.Endpoint, &buf)
	r.suite.NoError(err)

	// r.suite.T().Logf("%s %s", request.Method, request.URL.String())

	response, err := client.Do(request)
	r.suite.NoError(err)

	r.suite.Equal(rc, response.StatusCode, "Server responded with %s", response.Status)

	var responseBuf bytes.Buffer
	_, err = responseBuf.ReadFrom(response.Body)
	r.suite.NoError(err)

	expectedData, err := json.Marshal(data)
	r.suite.JSONEq(string(expectedData), responseBuf.String())
}
