var map = L.map('map');
const zoltanIcon = L.icon({
    iconUrl: '/icons/zoltan_icon.png',

    iconSize:     [38, 95], // size of the icon
    shadowSize:   [50, 64], // size of the shadow
    iconAnchor:   [22, 94], // point of the icon which will correspond to marker's location
    shadowAnchor: [4, 62],  // the same for the shadow
    popupAnchor:  [-3, -76] // point from which the popup should open relative to the iconAnchor
});

L.tileLayer('https://tile.openstreetmap.org/{z}/{x}/{y}.png', {
    attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
}).addTo(map);

const API = {
    getLocation: () => {
        return "46.9607673 18.9196372";
    }
}

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
            L.marker([location.latitude, location.longitude], {icon: zoltanIcon}).addTo(map)
                .bindPopup(k)
                .openPopup();
        });
    });
})();