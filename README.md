# bogie-pdm-frontend

This repository contains the frontend for the bogie PDM system in form of a Jupyter notebook.

The described setup is for Windows 11.

## Prepare:

* Clone this repository or download the zip file and extract it.
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
* Update pip: `python -m pip install --upgrade pip`

* Install Microsoft Visual C++ Build Tools (https://visualstudio.microsoft.com/visual-cpp-build-tools/): In the installation dialog, select "Desktop development with C++"


* Install Python dependencies:

```cmd
pip3 install -r analytics\requirements.txt
```

Ensure jupyter-lab is in the PATH. If not, add it manually, for example:

```cmd
set PATH=%PATH%;C:\Users\<usernam>\AppData\Local\Packages\PythonSoftwareFoundation.Python.3.10_qbz5n2kfra8p0\LocalCache\local-packages\Python310\Scripts
```

## Run:
```
jupyter-lab
```

Open `metricssimple.ipynb` in juptyer-lab. Run first two cells to load and display the metrics data.

Open `sensors_simple.ipynb` in juptyer-lab. Run first two cells to load and display the sensor data.




# Data format from stream

Data loaded from nats is converted from protobuf into a pandas dataframe. 

## Bogie dataframe

Contains the data from the vibration sensors.

Columns:
- `bogie_id`: Bogie ID (int). Can be used to identify the bogie if multiple bogies are monitored.
- `trigger_time`: Time when the trigger was detected (datetime)
- `seq`: Sequence number of this row (int). Can be used to uniquely identify the row. Sequence numbers are not necessarily sequential.
- `nats_rx_time`: Time when the data was received from NATS (datetime)
- `sensor_data`: dataframe with the sensor data. Columns:
    - `sensor0`: Data from sensor 0 (float)
    - `sensor1`: Data from sensor 1 (float)
    - `sensor2`: Data from sensor 2 (float)
    - `sensor3`: Data from sensor 3 (float)
- `lat`: Latitude (float) at the time of the trigger
- `lon`: Longitude (float) at the time of the trigger
- `alt`: Altitude (float) at the time of the trigger
- `speed`: Speed (float) at the time of the trigger
- `pos_valid`: True if the position data is valid (bool)

### Interpretation of Sensor Data

The sensor data in the bogie dataframe has a value range of -1 to 1. 

For voltage sensors, the following formula may be used to compute the voltage in V:

    voltage_in_volts = sensor_data * 10

For current sensors, the following formula may be used to compute the current in mA:

    current_in_ma = sensor_data * 20

When a 4..20mA acceleration sensor is used, the following formula may be used to compute the acceleration in g:

    acceleration_in_g = (sensor_data - 0.6) * (full_scale_g / 0.4)

where `full_scale_g` is the sensors full scale acceleration in g.

## Metrics dataframe

Contains the periodically published metrics, such as location, speed, LTE signal strength, etc.

Columns:
- `ts`: timestamp (datetime), when the metrics were published
- `seq`: Sequence number of this data frame (int). Can be used to uniquely identify the row. Sequence numbers are not necessarily sequential.
- `nats_rx_time`: Time when the data was received from NATS 
- `accel_x_rms`: RMS of the acceleration in the x direction (float)
- `accel_y_rms`: RMS of the acceleration in the y direction (float)
- `accel_z_rms`: RMS of the acceleration in the z direction (float)
- `accel_x_max`: Maximum acceleration in the x direction (float)
- `accel_y_max`: Maximum acceleration in the y direction (float)
- `accel_z_max`: Maximum acceleration in the z direction (float)
- `gnss_mode`: GNSS fix mode (int) 0=unknown, 1=no fix, 2=2D, 3=3D. 
These are only present if the `gnss_mode` is > 1:
    - `gnss_lat`: Latitude (float) 
    - `gnss_lon`: Longitude (float)
    - `gnss_alt`: Altitude (float)
    - `gnss_speed`: Speed (float)
    - `gnss_eph`: Estimated horizontal err in m (float)
- `gnss_num_sats`: Number of satellites used in the fix (int)
- `cellular_strength`: LTE signal strength in % (float)
- `cellular_operator`: LTE operator (string)
- `temperature_inbox`: Temperature in the Edge Computer (float)
- `internet_connected`: True if the Edge Computer is connected to the internet (bool)