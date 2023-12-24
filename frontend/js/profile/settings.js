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
    document.getElementById("profileButton").textContent = data['first_name'] + ' ' + data['last_name']
    document.getElementById("firstName").value = data['first_name']
    document.getElementById("lastName").value = data['last_name']
    if (data['public']) {
        document.getElementById("public").checked = true
    } else {
        document.getElementById("private").checked = true
    }
}

async function update() {
    let data = {
        first_name: document.getElementById("firstName").value,
        last_name: document.getElementById("lastName").value,
        public: document.getElementById("public").checked
    }
    await UpdateProfileSettingsRequest(data)
}