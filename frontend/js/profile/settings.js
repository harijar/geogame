import {GetProfileSettingsRequest, UpdateProfileSettingsRequest} from "../api.js";

window.onload = async function() {
    document.getElementById("mainMenuButton").onclick = async function() {
        window.location.href = '/index.html'
    }
    document.getElementById("backButton").onclick = async function() {
        window.location.href = '/profile'
    }
    document.getElementById("profileButton").onclick = async function() {
        window.location.href = '/profile'
    }
    await getInfo()
    document.getElementById("settings").onsubmit = update
}

async function getInfo() {
    let data = await GetProfileSettingsRequest()
    document.getElementById("profileButton").textContent = data['nickname']
    document.getElementById("nickname").value = data['nickname']
    if (data['public']) {
        document.getElementById("public").checked = true
    } else {
        document.getElementById("private").checked = true
    }
}

async function update() {
    let data = {
        nickname: document.getElementById("nickname").value,
        public: document.getElementById("public").checked
    }
    let error = await UpdateProfileSettingsRequest(data)
    if (error != null) {
        alert(error['error'])
    }
}