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

const displayModal = function (tID, endpoint) {
	console.log('diaplsy modal');
};

// init
insertDefaultDate(dateEl);
fillRandomValues();

// table.addEventListener('dblclick', function (e) {
// 	tRow = e.target.closest('.t');
// 	console.log(tRow);
// });
