import {AuthRequest, CheckAuthRequest} from "./api.js"

export async function Auth(user) {
    try {
        await AuthRequest(user)
    } catch (error) {
        alert(error)
        return
    }
    alert('Logged in as ' + user.username)
    window.location.href = '/profile/settings'
    alert('You can set up your profile here')
}

export async function CheckAuth() {
    const data = await CheckAuthRequest()
    if (data === null) {
        document.getElementById("auth").style.display = 'inline'
    } else {
        document.getElementById("profile").style.display = 'inline'
        document.getElementById("profileButton").textContent = data['nickname']
    }
}