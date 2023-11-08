function submitForm() {
    let country = document.getElementById("country").nodeValue
    fetch('/game/guess', {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify({country: country})
    })
        .then(response => {return response.json()})
}