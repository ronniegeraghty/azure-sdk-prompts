from managed_identity_demo.cli import main


def test_dry_run_is_offline_safe(capsys):
    exit_code = main(["--identity", "system"])

    assert exit_code == 0
    assert "Dry run complete" in capsys.readouterr().out


def test_user_mode_reports_missing_client_id(capsys):
    exit_code = main(["--identity", "user"])

    assert exit_code == 2
    assert "requires MANAGED_IDENTITY_CLIENT_ID" in capsys.readouterr().err
