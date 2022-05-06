const form = document.getElementById('form');
const errInsert = document.querySelector('.error-insert');

// to keep the default date to "today"
const insertDefaultDate = function () {
	const date = document.getElementById('date');

	const newDate = new Date();
	const year = newDate.getFullYear();
	const month = String(newDate.getMonth() + 1).padStart(2, '0');
	const day = String(newDate.getDate()).padStart(2, '0');

	date.value = `${year}-${month}-${day}`;
};

const sendFormData = async function (form) {
	const formData = new FormData(form);
	let reqUrl = `/add?`;

	Array.from(formData.keys()).forEach(key => {
		reqUrl += `${key}=${formData.get(key)}&`;
	});

	reqUrl = reqUrl.slice(0, -1);
	const res = await fetch(reqUrl);
	const resJson = await res.json();

	if (resJson.success) window.location.reload();
	else {
		errInsert.classList.remove('hidden');
		setTimeout(() => errInsert.classList.add('hidden'), 2500);
	}
};

const fillRandomValues = function () {
	const desc = (Math.random() + 1).toString(36).substring(2);
	const amount = Math.floor(Math.random() * (550 - 10) + 10);
	const paidTo = (Math.random() + 1).toString(36).substring(2);
	const mode = Math.random() > 0.5 ? true : false;
	const type = Math.random() > 0.5 ? true : false;

	document.getElementById('desc').value = desc;
	document.getElementById('amount').value = amount;
	document.getElementById('paid_to').value = paidTo;
	mode
		? (document.getElementById('mode').value = 'PhonePe')
		: (document.getElementById('mode').value = 'Google Pay');
	type
		? (document.getElementById('income').checked = true)
		: (document.getElementById('expense').checked = true);
};

insertDefaultDate();
fillRandomValues();

// event listeners
form.addEventListener('submit', e => {
	e.preventDefault();
	sendFormData(form);
});
