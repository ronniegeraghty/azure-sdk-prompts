from environment_detector import RuntimeEnvironment, detect_environment


def test_defaults_to_development() -> None:
    assert detect_environment({}) is RuntimeEnvironment.DEV


def test_detects_ci_before_other_markers() -> None:
    variables = {"GITHUB_WORKSPACE": "D:\\a\\repo", "IDENTITY_ENDPOINT": "http://host"}
    assert detect_environment(variables) is RuntimeEnvironment.CI


def test_detects_production_identity_host() -> None:
    assert (
        detect_environment({"AZURE_FEDERATED_TOKEN_FILE": "token.txt"})
        is RuntimeEnvironment.PRODUCTION
    )
