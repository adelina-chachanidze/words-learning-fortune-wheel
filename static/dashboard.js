// Load and display wheels
async function loadWheels() {
    try {
        const response = await fetch('/api/wheels');
        const wheels = await response.json();
        
        const wheelsList = document.getElementById('wheels-list');
        
        if (wheels.length === 0) {
            wheelsList.innerHTML = '<p>No wheels created yet.</p>';
            return;
        }
        
        let html = '';
        wheels.forEach(wheel => {
            if (wheel.length >= 3) {
                html += `
                    <div>
                        <h3>${wheel[1]}</h3>
                        <p>Wheel ID: ${wheel[0]} | Created: ${wheel[2]}</p>
                        <p>Student URL: /wheel/${wheel[0]}</p>
                        <hr>
                    </div>
                `;
            }
        });
        
        wheelsList.innerHTML = html;
    } catch (error) {
        document.getElementById('wheels-list').innerHTML = '<p>Error loading wheels.</p>';
    }
}

// Load wheels when page loads
document.addEventListener('DOMContentLoaded', loadWheels);