const daftar = document.getElementById('daftar');
const dataTersimpan = localStorage.getItem('riwayat');
if ( dataTersimpan=== null ) {
	daftar.textContent = "No history yet";
} else {
	const riwayat = JSON.parse(dataTersimpan);
	
	riwayat.reverse();
	const terbaru = riwayat.slice(0, 10);
	
	terbaru.forEach((item) => {
	const linkCode = document.createElement('a');
	const teksUrl = document.createElement('p');
	const btnDelete = document.createElement('button');
	btnDelete.textContent = 'Delete';
	
	btnDelete.addEventListener('click', function() {
		const dataTersimpan = localStorage.getItem('riwayat');
		const riwayat = JSON.parse(dataTersimpan);
		
		const riwayatBaru = riwayat.filter((x) => x.code !== item.code);
		localStorage.setItem('riwayat', JSON.stringify(riwayatBaru));
		newDiv.remove();
	});
	
	linkCode.textContent = item.code;
	linkCode.href = '/s/' + item.code;
	teksUrl.textContent = item.url;
	const newDiv = document.createElement('div')
	
	btnDelete.classList.add('btn-delete');
	newDiv.classList.add('riwayat-item')
	linkCode.classList.add('kotak-link');
	teksUrl.classList.add('teks-url');
	
	newDiv.append(linkCode, teksUrl, btnDelete);
	daftar.append(newDiv);
})};