import { LocalStorage } from "quasar";

var deviceID
function getDeviceID() {
    deviceID = LocalStorage.getItem("device_id");
    if (!deviceID) {
        deviceID = uuidv4()
        LocalStorage.set("device_id", deviceID);
    }

    return deviceID;
}

function uuidv4() {
    return ([1e7] + -1e3 + -4e3 + -8e3 + -1e11).replace(/[018]/g, c =>
        (
            c ^
            (crypto.getRandomValues(new Uint8Array(1))[0] & (15 >> (c / 4)))
        ).toString(16)
    );
}

function setPhoneNumber(number) {
    LocalStorage.set("phone_number", number);
}

function getPhoneNumber() {
    return LocalStorage.getItem("phone_number");
}

function setToken(token) {
    LocalStorage.set("token", token);
}

function getToken() {
    return LocalStorage.getItem("token");
}

export { getDeviceID, setPhoneNumber, getPhoneNumber, setToken, getToken }