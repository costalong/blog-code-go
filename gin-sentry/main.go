package main

// https://costa92.sentry.io/projects/
import (
	"fmt"
	"net/http"

	sentry "github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

func main() {
	// To initialize Sentry's handler, you need to initialize Sentry itself beforehand
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:           "https://807faac3cf2e42b15db2e49ab7d226c6@o84114.ingest.us.sentry.io/4507757794557952",
		EnableTracing: true,
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for tracing.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}

	// Then create your app
	app := gin.Default()

	// Once it's done, you can attach the handler as one of your middleware
	app.Use(sentrygin.New(sentrygin.Options{}))

	// Set up routes
	app.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello world!")
	})

	app.GET("/test", func(ctx *gin.Context) {
		// This is the line that sends the event to Sentry
		sentry.CaptureMessage("This is a test message!")
		ctx.String(http.StatusOK, "Test message sent!")
	})

	// And run it
	app.Run(":3000")

}
