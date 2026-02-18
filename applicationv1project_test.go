// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/CaseMark/casedev-go"
	"github.com/CaseMark/casedev-go/internal/testutil"
	"github.com/CaseMark/casedev-go/option"
)

func TestApplicationV1ProjectNewWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomcasemarkcasedevgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Applications.V1.Projects.New(context.TODO(), githubcomcasemarkcasedevgo.ApplicationV1ProjectNewParams{
		GitRepo:      githubcomcasemarkcasedevgo.F("gitRepo"),
		Name:         githubcomcasemarkcasedevgo.F("name"),
		BuildCommand: githubcomcasemarkcasedevgo.F("buildCommand"),
		EnvironmentVariables: githubcomcasemarkcasedevgo.F([]githubcomcasemarkcasedevgo.ApplicationV1ProjectNewParamsEnvironmentVariable{{
			Key:    githubcomcasemarkcasedevgo.F("key"),
			Target: githubcomcasemarkcasedevgo.F([]githubcomcasemarkcasedevgo.ApplicationV1ProjectNewParamsEnvironmentVariablesTarget{githubcomcasemarkcasedevgo.ApplicationV1ProjectNewParamsEnvironmentVariablesTargetProduction}),
			Value:  githubcomcasemarkcasedevgo.F("value"),
			Type:   githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.ApplicationV1ProjectNewParamsEnvironmentVariablesTypePlain),
		}}),
		Framework:       githubcomcasemarkcasedevgo.F("framework"),
		GitBranch:       githubcomcasemarkcasedevgo.F("gitBranch"),
		InstallCommand:  githubcomcasemarkcasedevgo.F("installCommand"),
		OutputDirectory: githubcomcasemarkcasedevgo.F("outputDirectory"),
		RootDirectory:   githubcomcasemarkcasedevgo.F("rootDirectory"),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestApplicationV1ProjectGet(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomcasemarkcasedevgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Applications.V1.Projects.Get(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestApplicationV1ProjectList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomcasemarkcasedevgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Applications.V1.Projects.List(context.TODO())
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestApplicationV1ProjectDeleteWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomcasemarkcasedevgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Applications.V1.Projects.Delete(
		context.TODO(),
		"id",
		githubcomcasemarkcasedevgo.ApplicationV1ProjectDeleteParams{
			DeleteFromHosting: githubcomcasemarkcasedevgo.F(true),
		},
	)
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestApplicationV1ProjectNewDeploymentWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomcasemarkcasedevgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Applications.V1.Projects.NewDeployment(
		context.TODO(),
		"id",
		githubcomcasemarkcasedevgo.ApplicationV1ProjectNewDeploymentParams{
			EnvironmentVariables: githubcomcasemarkcasedevgo.F([]githubcomcasemarkcasedevgo.ApplicationV1ProjectNewDeploymentParamsEnvironmentVariable{{
				Key:    githubcomcasemarkcasedevgo.F("key"),
				Target: githubcomcasemarkcasedevgo.F([]githubcomcasemarkcasedevgo.ApplicationV1ProjectNewDeploymentParamsEnvironmentVariablesTarget{githubcomcasemarkcasedevgo.ApplicationV1ProjectNewDeploymentParamsEnvironmentVariablesTargetProduction}),
				Value:  githubcomcasemarkcasedevgo.F("value"),
				Type:   githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.ApplicationV1ProjectNewDeploymentParamsEnvironmentVariablesTypePlain),
			}}),
		},
	)
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestApplicationV1ProjectNewDomainWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomcasemarkcasedevgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Applications.V1.Projects.NewDomain(
		context.TODO(),
		"id",
		githubcomcasemarkcasedevgo.ApplicationV1ProjectNewDomainParams{
			Domain:    githubcomcasemarkcasedevgo.F("domain"),
			GitBranch: githubcomcasemarkcasedevgo.F("gitBranch"),
		},
	)
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestApplicationV1ProjectNewEnvWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomcasemarkcasedevgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Applications.V1.Projects.NewEnv(
		context.TODO(),
		"id",
		githubcomcasemarkcasedevgo.ApplicationV1ProjectNewEnvParams{
			Key:       githubcomcasemarkcasedevgo.F("key"),
			Target:    githubcomcasemarkcasedevgo.F([]githubcomcasemarkcasedevgo.ApplicationV1ProjectNewEnvParamsTarget{githubcomcasemarkcasedevgo.ApplicationV1ProjectNewEnvParamsTargetProduction}),
			Value:     githubcomcasemarkcasedevgo.F("value"),
			GitBranch: githubcomcasemarkcasedevgo.F("gitBranch"),
			Type:      githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.ApplicationV1ProjectNewEnvParamsTypePlain),
		},
	)
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestApplicationV1ProjectDeleteDomain(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomcasemarkcasedevgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Applications.V1.Projects.DeleteDomain(
		context.TODO(),
		"id",
		"domain",
	)
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestApplicationV1ProjectDeleteEnv(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomcasemarkcasedevgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Applications.V1.Projects.DeleteEnv(
		context.TODO(),
		"id",
		"envId",
	)
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestApplicationV1ProjectGetRuntimeLogsWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomcasemarkcasedevgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Applications.V1.Projects.GetRuntimeLogs(
		context.TODO(),
		"id",
		githubcomcasemarkcasedevgo.ApplicationV1ProjectGetRuntimeLogsParams{
			Limit: githubcomcasemarkcasedevgo.F(0.000000),
		},
	)
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestApplicationV1ProjectListDeploymentsWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomcasemarkcasedevgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Applications.V1.Projects.ListDeployments(
		context.TODO(),
		"id",
		githubcomcasemarkcasedevgo.ApplicationV1ProjectListDeploymentsParams{
			Limit:  githubcomcasemarkcasedevgo.F(0.000000),
			State:  githubcomcasemarkcasedevgo.F("state"),
			Target: githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.ApplicationV1ProjectListDeploymentsParamsTargetProduction),
		},
	)
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestApplicationV1ProjectListDomains(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomcasemarkcasedevgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Applications.V1.Projects.ListDomains(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestApplicationV1ProjectListEnvWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomcasemarkcasedevgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Applications.V1.Projects.ListEnv(
		context.TODO(),
		"id",
		githubcomcasemarkcasedevgo.ApplicationV1ProjectListEnvParams{
			Decrypt: githubcomcasemarkcasedevgo.F(true),
		},
	)
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
