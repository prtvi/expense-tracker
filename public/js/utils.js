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
const DEL_ENDPOINT = '/del';

// keys for session storage
const UPDATE_TID = 'UPDATE_TID';
const UPDATE_TRUE = 'UPDATE';

//
//

/*
	UPDATE process
	- event listener for the table checks if edit icon is clicked
	- get the transaction id 
	- fetch the document (transaction) from the db and display for edit
	- store this document id in sessionStorage to fetch after reload
	- look for form submission, if the btn text is 'Update transaction' then send update_id and data and update the transaction
	- use the stored tID to render the update
*/

//
//

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

// on success
const onSuccess = function () {
	window.location.reload();
};

// on form submission failure
const onFailure = function (endpoint) {
	if (endpoint === ADD_ENDPOINT) showError('Error inserting the transaction!');

	if (endpoint === EDIT_ENDPOINT) showError('Error updating the transaction!');
};

const highLightT = function (tID) {
	const tRow = document.getElementById(tID);

	setTimeout(() => {
		setTimeout(() => {
			setTimeout(() => {
				setTimeout(() => {
					tRow.classList.toggle('updated-t');
				}, 300);
				tRow.classList.toggle('updated-t');
			}, 300);
			tRow.classList.toggle('updated-t');
		}, 300);
		tRow.classList.toggle('updated-t');
	}, 300);
};

// init
insertDefaultDate(dateEl);
fillRandomValues();
