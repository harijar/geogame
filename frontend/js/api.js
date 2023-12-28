const API_URL = 'http://localhost:8080/v1/'
const GAME_START_URL = API_URL + 'game/start'
const GAME_GUESS_URL = API_URL + 'game/guess'
const AUTH_URL = API_URL + 'auth'
const PROFILE_URL = API_URL + 'profile'
const PROFILE_SETTINGS_URL = API_URL + 'profile/settings'

export async function StartGameRequest() {
    const response= await fetch(GAME_START_URL, {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        credentials: 'include',
    })
    if (response.ok) {
        const data = await response.json()
        if (!data) {
            throw Error('Failed to start: ' + await response.text())
        }
        return data
    }
    throw Error('Failed to start: ' + response.status)
}

export async function GuessGameRequest(guess) {
    const response = await fetch(GAME_GUESS_URL, {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        credentials: 'include',
        body: JSON.stringify({guess: guess})
    })
    if (response.ok) {
        const data = await response.json()
        if (!data) {
            throw Error('Failed to guess: ' + await response.text())
        }
        return data
    }
    if (response.status === 404) { // game hasn't started
        return null
    }
    throw Error('Failed to guess: ' + response.status)
}

export async function AuthRequest(user) {
    const response = await fetch(AUTH_URL, {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        credentials: 'include',
        body: JSON.stringify(user)
    })
    if (!response.ok) {
        throw Error("Failed to authorize: " + response.status)
    }
}

export async function CheckAuthRequest() {
    const response = await fetch(AUTH_URL, {
        method: 'GET',
        credentials: 'include'
    })
    if (response.ok) {
        const data = await response.json();
        if (!data) {
            throw Error('Failed to check auth: ' + await response.text())
        }
        return data
    }
    if (response.status === 403) {
        return null
    }
    throw Error('Failed to check auth: ' + response.status)
}

export async function ProfileRequest() {
    const response = await fetch(PROFILE_URL, {
        method: 'GET',
        credentials: 'include'
    })
    if (response.ok) {
        const data = await response.json()
        if (!data) {
            throw Error('Failed to get profile info: ' + await response.text())
        }
        return data
    }
    if (response.status === 403) {
        return null
    }
    throw Error('Failed to get profile info: ' + response.status)
}

export async function GetProfileSettingsRequest() {
    const response = await fetch(PROFILE_SETTINGS_URL, {
        method: 'GET',
        credentials: 'include'
    })
    if (response.ok) {
        const data = await response.json()
        if (!data) {
            throw Error('Failed to get profile info: ' + await response.text())
        }
        return data
    }
    throw Error('Failed to get profile info: ' + response.status)
}

export async function UpdateProfileSettingsRequest(settings) {
    const response = await fetch(PROFILE_SETTINGS_URL, {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        credentials: 'include',
        body: JSON.stringify(settings)
    })
    if (!response.ok) {
        if (response.status < 500) {
            const data = await response.json()
            if (!data) {
                throw Error('Failed to update: ' + await response.text())
            }
            return data
        }
        throw Error('Failed to update: ' + response.status)
    }
    return null
}