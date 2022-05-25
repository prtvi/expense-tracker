/**
 * To keep the default date to "today"
 * @param {object} dateEl Date - input form element
 */

const insertDefaultDate = function (dateEl) {
	const newDate = new Date();
	const year = newDate.getFullYear();
	const month = String(newDate.getMonth() + 1).padStart(2, '0');
	const day = String(newDate.getDate()).padStart(2, '0');

	dateEl.value = `${year}-${month}-${day}`;
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

// renders error popup with given "errMsg"

/**
 * Render error poput with given "errMsg"
 * @param {string} errMsg Error message to display
 */

const showError = function (errMsg) {
	errText.textContent = errMsg;
	errDiv.classList.remove('hidden');
	setTimeout(() => errDiv.classList.add('hidden'), 2500);
};

// on success event
const onSuccess = () => window.location.reload();

/**
 * Render error on form submission failure
 * @param {string} endpoint Error based on endpoint
 */
const onFailure = function (endpoint) {
	if (endpoint === ADD_ENDPOINT) showError('Error inserting the transaction!');
	if (endpoint === EDIT_ENDPOINT) showError('Error updating the transaction!');
};

const displayTModal = async function (tID, endpoint) {
	const url = `${endpoint}?id=${tID}`;
	const res = await makeFetchRequest(url);

	if (!res.date) return showError('Error loading the transaction!');

	let formattedDate = String(new Date(res.date)).slice(0, 16);
	formattedDate = `${formattedDate.slice(0, 3)}, ${formattedDate.slice(4, -1)}`;

	// TODO: add currency symbol here
	const map = new Map([
		['Date', formattedDate],
		['Description', res.desc],
		['Amount', res.amount],
		['Mode', res.mode],
		['Paid to', res.paid_to],
		['Type', res.type],
	]);

	let fieldContainers = '';
	map.forEach((value, key) => {
		fieldContainers += `<div class="modal-t-field-container">
	<label class="modal-t-field">${key}</label>
	<p class="modal-t-field-value">${value}</p>
	</div>`;
	});

	modalContent.innerHTML =
		`<h3 class="modal-title">${res.desc} on ${formattedDate}</h3>` +
		fieldContainers;

	showModal();
};

// init
insertDefaultDate(dateEl);
fillRandomValues();

// table.addEventListener('dblclick', function (e) {
// 	tRow = e.target.closest('.t');
// 	console.log(tRow);
// });
