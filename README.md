# bogie-pdm-frontend

## Prepare:

* Check if the settings in `analysis/account.py` are correct.
* Create `analysis/user.creds` file with NATS user JWT and NKEY seed. Example:

```
-----BEGIN NATS USER JWT-----
eyJ0eX....nbSAMcQDq_Qm0Iw5qX3V-mxRslBLox85Bt7hZf5ZtwGTWHPVYZB0av3F7TWVU0MXYxqACg
------END NATS USER JWT------

************************* IMPORTANT *************************
NKEY Seed printed below can be used to sign and prove identity.
NKEYs are sensitive and should be treated as secrets.

-----BEGIN USER NKEY SEED-----
SUAG3O2H6K....6K3SKLM7VPLPL73DJI
------END USER NKEY SEED------

*************************************************************
```

* Install Python (tested with 3.10), downloaded via MS AppStore

* `python -m pip install --upgrade pip`

* Install Microsoft Visual C++ Build Tools (https://visualstudio.microsoft.com/visual-cpp-build-tools/), and open command prompt by "Start" the tools.

* Install Python dependencies:

```bash
pip3 install -r analytics/requirements.txt
```

## Run:
```
jupyter-lab
```
