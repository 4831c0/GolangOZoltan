const h = window.screen.height;

for (let i = 0; i < (h/20); i++) {
    document.getElementById('map').innerText += '\n';
}

var map = L.map('map');

L.tileLayer('https://tile.openstreetmap.org/{z}/{x}/{y}.png', {
    attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
}).addTo(map);

function updateLocation() {
    const [ lat, long ] = API.getLocation().split(" ").map(v => parseFloat(v));
    // document.getElementById('debug').innerText = `${lat} ${long}`;
    map.setView([lat, long], 18);
}

setInterval(updateLocation, 100);

(async ()=>{
    const req = await fetch('/api/zoltan/list');
    const zoltanList = await req.json();

    Object.keys(zoltanList).forEach(k => {
        const zoltan = zoltanList[k];

        zoltan.locations.forEach(location => {
            L.marker([location.latitude, location.longitude]).addTo(map)
                .bindPopup(k)
                .openPopup();
        });
    });
})();