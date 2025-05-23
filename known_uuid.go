package gatt

// A dictionary of known service names and type (keyed by service uuid)
var knownServices = map[string]struct{ Name, Type string }{
	"1800": {Name: "Generic Access", Type: "org.bluetooth.service.generic_access"},
	"1801": {Name: "Generic Attribute", Type: "org.bluetooth.service.generic_attribute"},
	"1802": {Name: "Immediate Alert", Type: "org.bluetooth.service.immediate_alert"},
	"1803": {Name: "Link Loss", Type: "org.bluetooth.service.link_loss"},
	"1804": {Name: "Tx Power", Type: "org.bluetooth.service.tx_power"},
	"1805": {Name: "Current Time Service", Type: "org.bluetooth.service.current_time"},
	"1806": {Name: "Reference Time Update Service", Type: "org.bluetooth.service.reference_time_update"},
	"1807": {Name: "Next DST Change Service", Type: "org.bluetooth.service.next_dst_change"},
	"1808": {Name: "Glucose", Type: "org.bluetooth.service.glucose"},
	"1809": {Name: "Health Thermometer", Type: "org.bluetooth.service.health_thermometer"},
	"180a": {Name: "Device Information", Type: "org.bluetooth.service.device_information"},
	"180d": {Name: "Heart Rate", Type: "org.bluetooth.service.heart_rate"},
	"180e": {Name: "Phone Alert Status Service", Type: "org.bluetooth.service.phone_alert_service"},
	"180f": {Name: "Battery Service", Type: "org.bluetooth.service.battery_service"},
	"1810": {Name: "Blood Pressure", Type: "org.bluetooth.service.blood_pressuer"},
	"1811": {Name: "Alert Notification Service", Type: "org.bluetooth.service.alert_notification"},
	"1812": {Name: "Human Interface Device", Type: "org.bluetooth.service.human_interface_device"},
	"1813": {Name: "Scan Parameters", Type: "org.bluetooth.service.scan_parameters"},
	"1814": {Name: "Running Speed and Cadence", Type: "org.bluetooth.service.running_speed_and_cadence"},
	"1815": {Name: "Cycling Speed and Cadence", Type: "org.bluetooth.service.cycling_speed_and_cadence"},
}

func RegisterService(uuid string, name string, tp string) {
	knownServices[uuid] = struct{ Name, Type string }{Name: name, Type: tp}
}

// A dictionary of known descriptor names and type (keyed by attribute uuid)
var knownAttributes = map[string]struct{ Name, Type string }{
	"2800": {Name: "Primary Service", Type: "org.bluetooth.attribute.gatt.primary_service_declaration"},
	"2801": {Name: "Secondary Service", Type: "org.bluetooth.attribute.gatt.secondary_service_declaration"},
	"2802": {Name: "Include", Type: "org.bluetooth.attribute.gatt.include_declaration"},
	"2803": {Name: "Characteristic", Type: "org.bluetooth.attribute.gatt.characteristic_declaration"},
}

func RegisterAttribute(uuid string, name string, tp string) {
	knownAttributes[uuid] = struct{ Name, Type string }{Name: name, Type: tp}
}

// A dictionary of known descriptor names and type (keyed by descriptor uuid)
var knownDescriptors = map[string]struct{ Name, Type string }{
	"2900": {Name: "Characteristic Extended Properties", Type: "org.bluetooth.descriptor.gatt.characteristic_extended_properties"},
	"2901": {Name: "Characteristic User Description", Type: "org.bluetooth.descriptor.gatt.characteristic_user_description"},
	"2902": {Name: "Client Characteristic Configuration", Type: "org.bluetooth.descriptor.gatt.client_characteristic_configuration"},
	"2903": {Name: "Server Characteristic Configuration", Type: "org.bluetooth.descriptor.gatt.server_characteristic_configuration"},
	"2904": {Name: "Characteristic Presentation Format", Type: "org.bluetooth.descriptor.gatt.characteristic_presentation_format"},
	"2905": {Name: "Characteristic Aggregate Format", Type: "org.bluetooth.descriptor.gatt.characteristic_aggregate_format"},
	"2906": {Name: "Valid Range", Type: "org.bluetooth.descriptor.valid_range"},
	"2907": {Name: "External Report Reference", Type: "org.bluetooth.descriptor.external_report_reference"},
	"2908": {Name: "Report Reference", Type: "org.bluetooth.descriptor.report_reference"},
}

func RegisterDescriptor(uuid string, name string, tp string) {
	knownDescriptors[uuid] = struct{ Name, Type string }{Name: name, Type: tp}
}

// A dictionary of known characteristic names and type (keyed by characteristic uuid)
var knownCharacteristics = map[string]struct{ Name, Type string }{
	"2a00": {Name: "Device Name", Type: "org.bluetooth.characteristic.gap.device_name"},
	"2a01": {Name: "Appearance", Type: "org.bluetooth.characteristic.gap.appearance"},
	"2a02": {Name: "Peripheral Privacy Flag", Type: "org.bluetooth.characteristic.gap.peripheral_privacy_flag"},
	"2a03": {Name: "Reconnection Address", Type: "org.bluetooth.characteristic.gap.reconnection_address"},
	"2a04": {Name: "Peripheral Preferred Connection Parameters", Type: "org.bluetooth.characteristic.gap.peripheral_preferred_connection_parameters"},
	"2a05": {Name: "Service Changed", Type: "org.bluetooth.characteristic.gatt.service_changed"},
	"2a06": {Name: "Alert Level", Type: "org.bluetooth.characteristic.alert_level"},
	"2a07": {Name: "Tx Power Level", Type: "org.bluetooth.characteristic.tx_power_level"},
	"2a08": {Name: "Date Time", Type: "org.bluetooth.characteristic.date_time"},
	"2a09": {Name: "Day of Week", Type: "org.bluetooth.characteristic.day_of_week"},
	"2a0a": {Name: "Day Date Time", Type: "org.bluetooth.characteristic.day_date_time"},
	"2a0c": {Name: "Exact Time 256", Type: "org.bluetooth.characteristic.exact_time_256"},
	"2a0d": {Name: "DST Offset", Type: "org.bluetooth.characteristic.dst_offset"},
	"2a0e": {Name: "Time Zone", Type: "org.bluetooth.characteristic.time_zone"},
	"2a0f": {Name: "Local Time Information", Type: "org.bluetooth.characteristic.local_time_information"},
	"2a11": {Name: "Time with DST", Type: "org.bluetooth.characteristic.time_with_dst"},
	"2a12": {Name: "Time Accuracy", Type: "org.bluetooth.characteristic.time_accuracy"},
	"2a13": {Name: "Time Source", Type: "org.bluetooth.characteristic.time_source"},
	"2a14": {Name: "Reference Time Information", Type: "org.bluetooth.characteristic.reference_time_information"},
	"2a16": {Name: "Time Update Control Point", Type: "org.bluetooth.characteristic.time_update_control_point"},
	"2a17": {Name: "Time Update State", Type: "org.bluetooth.characteristic.time_update_state"},
	"2a18": {Name: "Glucose Measurement", Type: "org.bluetooth.characteristic.glucose_measurement"},
	"2a19": {Name: "Battery Level", Type: "org.bluetooth.characteristic.battery_level"},
	"2a1c": {Name: "Temperature Measurement", Type: "org.bluetooth.characteristic.temperature_measurement"},
	"2a1d": {Name: "Temperature Type", Type: "org.bluetooth.characteristic.temperature_type"},
	"2a1e": {Name: "Intermediate Temperature", Type: "org.bluetooth.characteristic.intermediate_temperature"},
	"2a21": {Name: "Measurement Interval", Type: "org.bluetooth.characteristic.measurement_interval"},
	"2a22": {Name: "Boot Keyboard Input Report", Type: "org.bluetooth.characteristic.boot_keyboard_input_report"},
	"2a23": {Name: "System ID", Type: "org.bluetooth.characteristic.system_id"},
	"2a24": {Name: "Model Number String", Type: "org.bluetooth.characteristic.model_number_string"},
	"2a25": {Name: "Serial Number String", Type: "org.bluetooth.characteristic.serial_number_string"},
	"2a26": {Name: "Firmware Revision String", Type: "org.bluetooth.characteristic.firmware_revision_string"},
	"2a27": {Name: "Hardware Revision String", Type: "org.bluetooth.characteristic.hardware_revision_string"},
	"2a28": {Name: "Software Revision String", Type: "org.bluetooth.characteristic.software_revision_string"},
	"2a29": {Name: "Manufacturer Name String", Type: "org.bluetooth.characteristic.manufacturer_name_string"},
	"2a2a": {Name: "IEEE 11073-20601 Regulatory Certification Data List", Type: "org.bluetooth.characteristic.ieee_11073-20601_regulatory_certification_data_list"},
	"2a2b": {Name: "Current Time", Type: "org.bluetooth.characteristic.current_time"},
	"2a31": {Name: "Scan Refresh", Type: "org.bluetooth.characteristic.scan_refresh"},
	"2a32": {Name: "Boot Keyboard Output Report", Type: "org.bluetooth.characteristic.boot_keyboard_output_report"},
	"2a33": {Name: "Boot Mouse Input Report", Type: "org.bluetooth.characteristic.boot_mouse_input_report"},
	"2a34": {Name: "Glucose Measurement Context", Type: "org.bluetooth.characteristic.glucose_measurement_context"},
	"2a35": {Name: "Blood Pressure Measurement", Type: "org.bluetooth.characteristic.blood_pressure_measurement"},
	"2a36": {Name: "Intermediate Cuff Pressure", Type: "org.bluetooth.characteristic.intermediate_blood_pressure"},
	"2a37": {Name: "Heart Rate Measurement", Type: "org.bluetooth.characteristic.heart_rate_measurement"},
	"2a38": {Name: "Body Sensor Location", Type: "org.bluetooth.characteristic.body_sensor_location"},
	"2a39": {Name: "Heart Rate Control Point", Type: "org.bluetooth.characteristic.heart_rate_control_point"},
	"2a3f": {Name: "Alert Status", Type: "org.bluetooth.characteristic.alert_status"},
	"2a40": {Name: "Ringer Control Point", Type: "org.bluetooth.characteristic.ringer_control_point"},
	"2a41": {Name: "Ringer Setting", Type: "org.bluetooth.characteristic.ringer_setting"},
	"2a42": {Name: "Alert Category ID Bit Mask", Type: "org.bluetooth.characteristic.alert_category_id_bit_mask"},
	"2a43": {Name: "Alert Category ID", Type: "org.bluetooth.characteristic.alert_category_id"},
	"2a44": {Name: "Alert Notification Control Point", Type: "org.bluetooth.characteristic.alert_notification_control_point"},
	"2a45": {Name: "Unread Alert Status", Type: "org.bluetooth.characteristic.unread_alert_status"},
	"2a46": {Name: "New Alert", Type: "org.bluetooth.characteristic.new_alert"},
	"2a47": {Name: "Supported New Alert Category", Type: "org.bluetooth.characteristic.supported_new_alert_category"},
	"2a48": {Name: "Supported Unread Alert Category", Type: "org.bluetooth.characteristic.supported_unread_alert_category"},
	"2a49": {Name: "Blood Pressure Feature", Type: "org.bluetooth.characteristic.blood_pressure_feature"},
	"2a4a": {Name: "HID Information", Type: "org.bluetooth.characteristic.hid_information"},
	"2a4b": {Name: "Report Map", Type: "org.bluetooth.characteristic.report_map"},
	"2a4c": {Name: "HID Control Point", Type: "org.bluetooth.characteristic.hid_control_point"},
	"2a4d": {Name: "Report", Type: "org.bluetooth.characteristic.report"},
	"2a4e": {Name: "Protocol Mode", Type: "org.bluetooth.characteristic.protocol_mode"},
	"2a4f": {Name: "Scan Interval Window", Type: "org.bluetooth.characteristic.scan_interval_window"},
	"2a50": {Name: "PnP ID", Type: "org.bluetooth.characteristic.pnp_id"},
	"2a51": {Name: "Glucose Feature", Type: "org.bluetooth.characteristic.glucose_feature"},
	"2a52": {Name: "Record Access Control Point", Type: "org.bluetooth.characteristic.record_access_control_point"},
	"2a53": {Name: "RSC Measurement", Type: "org.bluetooth.characteristic.rsc_measurement"},
	"2a54": {Name: "RSC Feature", Type: "org.bluetooth.characteristic.rsc_feature"},
	"2a55": {Name: "SC Control Point", Type: "org.bluetooth.characteristic.sc_control_point"},
	"2a5b": {Name: "CSC Measurement", Type: "org.bluetooth.characteristic.csc_measurement"},
	"2a5c": {Name: "CSC Feature", Type: "org.bluetooth.characteristic.csc_feature"},
	"2a5d": {Name: "Sensor Location", Type: "org.bluetooth.characteristic.sensor_location"},
}

func RegisterCharacteristic(uuid string, name string, tp string) {
	knownCharacteristics[uuid] = struct{ Name, Type string }{Name: name, Type: tp}
}
