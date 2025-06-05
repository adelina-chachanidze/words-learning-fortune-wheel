let words = [];
let currentWord = '';

// Get wheel ID from URL
function getWheelID() {
    return window.location.pathname.split('/wheel/')[1];
}

// Load words for this wheel
async function loadWords() {
    const wheelID = getWheelID();
    
    try {
        const response = await fetch(`/api/words/${wheelID}`);
        words = await response.json();
        
        if (words.length === 0) {
            document.getElementById('current-word').textContent = 'No words found!';
            document.getElementById('progress').textContent = 'Wheel not found or empty.';
        } else {
            document.getElementById('progress').textContent = `${words.length} words available`;
        }
    } catch (error) {
        document.getElementById('current-word').textContent = 'Error loading words!';
    }
}

function spinWheel() {
    if (words.length === 0) {
        document.getElementById('current-word').textContent = 'No words available';
        return;
    }
    
    const randomWord = words[Math.floor(Math.random() * words.length)];
    currentWord = randomWord;
    document.getElementById('current-word').textContent = randomWord;
}

function rememberWord() {
    alert('Great! You remembered: ' + currentWord);
    // Later: send to server to update CSV
}

function forgetWord() {
    alert('No worries! Keep practicing: ' + currentWord);
    // Later: send to server to update CSV
}

// Load words when page loads
document.addEventListener('DOMContentLoaded', loadWords);