const daftar = document.getElementById('daftar');
const dataTersimpan = localStorage.getItem('riwayat');
if ( dataTersimpan=== null ) {
	daftar.textContent = "No history yet";
} else {
	const riwayat = JSON.parse(dataTersimpan);
	
	riwayat.reverse();
	const terbaru = riwayat.slice(0, 10);
	
	terbaru.forEach((item) => {
	const card = document.createElement('div');
	const linkCode = document.createElement('a');
	const teksUrl = document.createElement('p');
	
	linkCode.textContent = item.code;
	linkCode.href = '/s/' + item.code;
	teksUrl.textContent = item.url;
	
	card.append(linkCode);
	card.append(teksUrl);
	daftar.append(card);
})};