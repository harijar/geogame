const API_URL = 'http://localhost:8080/v1/';
const GAME_START_URL = API_URL + 'game/start';
const GAME_GUESS_URL = API_URL + 'game/guess';

export async function StartGameRequest() {
    const response= await fetch(GAME_START_URL, {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        credentials: 'include',
    });
    if (response.ok) {
        const data = await response.json();
        if (!data) {
            throw Error('Failed to start: ' + await response.text());
        }
        return data
    }
    throw Error('Failed to start: ' + response.status);
}

export async function GuessGameRequest(guess) {
    const response = await fetch(GAME_GUESS_URL, {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        credentials: 'include',
        body: JSON.stringify({guess: guess})
    });
    if (response.ok) {
        const data = await response.json();
        if (!data) {
            throw Error('Failed to guess: ' + await response.text());
        }
        return data;
    }
    if (response.status === 404) { // game hasn't started
        return null;
    }
    throw Error('Failed to guess: ' + response.status);
}