const form = document.getElementById('form');
const errDiv = document.querySelector('.error-div');
const errText = document.querySelector('.error-text');
const table = document.querySelector('.t-table');

// form elements, except type
const dateEl = document.getElementById('date');
const descEl = document.getElementById('desc');
const amountEl = document.getElementById('amount');
const paidToEl = document.getElementById('paid_to');
const modeEl = document.getElementById('mode');

// btns
const submitBtn = document.querySelector('.btn-add');

// endpoints
const GET_ENDPOINT = '/get';
const ADD_ENDPOINT = '/add';
const EDIT_ENDPOINT = '/edit';

// to keep the default date to "today"
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

// make fetch request and get back json response
const makeFetchRequest = async function (url) {
	const res = await fetch(url);
	const resJson = await res.json();

	return resJson;
};

// renders error popup with given "errMsg"
const showError = function (errMsg) {
	errText.textContent = errMsg;
	errDiv.classList.remove('hidden');
	setTimeout(() => errDiv.classList.add('hidden'), 2500);
};

// send current form data to backend
// if update === false, then sends form data for insert/add operation
// if update === true, then send UPDATE_TID (global variable) to update the loaded transaction
const sendFormData = async function (
	form,
	endpoint,
	update = false,
	UPDATE_TID = null
) {
	const formData = new FormData(form);
	let reqUrl = `${endpoint}?`;

	// generate query with form data
	Array.from(formData.keys()).forEach(key => {
		reqUrl += `${key}=${formData.get(key)}&`;
	});

	// if update then append t_id to query
	if (update) reqUrl += `id=${UPDATE_TID}`;
	// else remove the last '&' from query
	else reqUrl = reqUrl.slice(0, -1);

	const res = await makeFetchRequest(reqUrl);

	// if success then reload window to reload transactions
	if (res.success) window.location.reload();
	// else render errors accordingly
	else {
		if (endpoint === ADD_ENDPOINT)
			showError('Error inserting the transaction!');

		if (endpoint === EDIT_ENDPOINT)
			showError('Error updating the transaction!');
	}
};

// global variable to get tID when update is clicked
let UPDATE_TID;

// get the transaction data from the backend, insert it into form to edit, init UPDATE_TID
const getAndLoadTForEdit = async function (tID, endpoint) {
	const url = `${endpoint}?id=${tID}`;
	const res = await makeFetchRequest(url);
	if (!res.date) return showError('Error loading the transaction!');

	// load values to ui
	dateEl.value = res.date;
	descEl.value = res.desc;
	amountEl.value = res.amount;
	paidToEl.value = res.paid_to;
	modeEl.value = res.mode;

	if (res.type === 'Income') document.getElementById('income').checked = true;
	else document.getElementById('expense').checked = true;

	// change btn text content
	submitBtn.textContent = 'Update transaction';

	// init UPDATE_TID
	UPDATE_TID = res._id;
};

// using event delegation to add event listener to the entire table rather than every transaction
table &&
	table.addEventListener('click', e => {
		const tRow = e.target.closest('.t');
		const tID = tRow.getAttribute('t_id');

		// looking for edit event
		if (e.target.classList.contains('edit-icon'))
			getAndLoadTForEdit(tID, GET_ENDPOINT);

		// delete event
		if (e.target.classList.contains('delete-icon'))
			console.log('delete icon clicked');
	});

form.addEventListener('submit', e => {
	e.preventDefault();

	// if btn text-content is for adding transaction then send form data wo update options
	if (submitBtn.textContent === 'Add transaction')
		sendFormData(form, ADD_ENDPOINT, false);
	// else send UPDATE_TID
	else sendFormData(form, EDIT_ENDPOINT, true, UPDATE_TID);
});

insertDefaultDate(dateEl);
fillRandomValues();
