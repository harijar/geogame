function startGame() {
    fetch('/game/start', {
        method: 'POST'
    })
        .then(response => response.json())
    window.location.href = '/game/guess'
}