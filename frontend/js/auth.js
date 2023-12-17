import {AuthRequest, CheckAuthRequest} from "./api.js"

window.auth = auth

export async function Auth(user) {
    try {
        await AuthRequest(user)
    } catch (error) {
        alert(error)
        return
    }
    alert('Logged in as ' + user.first_name + ' ' + user.last_name + ' (' + user.id + (user.username ? ', @' + user.username : '') + ')')
}

export async function CheckAuth() {
    const data = await CheckAuthRequest()
    if (data === null) {
        document.getElementById("auth").style.display = 'inline'
    } else {
        document.getElementById("profile").style.display = 'inline'
        document.getElementById("profileButton").textContent = data['first_name'] + ' ' + data['last_name']
    }
}