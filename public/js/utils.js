// util functions

/**
 * Make fetch request and get back JSON response
 *
 * @param {string} url URL / endpoint to which request will be made
 * @returns {JSON} JSON response from server
 */

const makeFetchRequest = async function (url) {
	const res = await fetch(url);
	const resJson = await res.json();
	return resJson;
};

/**
 * Send current form data to backend
 *
 * @param {object} form Entire html form element
 * @param {string} endpoint Endpoint to which request will be made
 * @param {boolean} update Check if update or edit request is to be made
 * @param {string} tID Transaction ID of the transaction to be updated
 *
 * if update === false, then sends form data for insert transaction operation
 * if update === true, then sends tID to update the updated transaction
 */

const sendFormData = async function (
	form,
	endpoint,
	update = false,
	tID = null
) {
	const formData = new FormData(form);
	let reqUrl = `${endpoint}?`;

	// generate query with form data
	Array.from(formData.keys()).forEach(
		key => (reqUrl += `${key}=${formData.get(key)}&`)
	);

	// if update then append id to query
	if (update) reqUrl += `id=${tID}`;
	// else remove the last '&' from query
	else reqUrl = reqUrl.slice(0, -1);

	const res = await makeFetchRequest(reqUrl);

	// if success then reload window to reload transactions
	if (res.success) onSuccess();
	// else render errors accordingly
	else onFailure(endpoint);
};

/**
 * Get the transaction data from the backend, insert it into form to edit, init UPDATE_ID to reflect update on reload
 *
 * @param {string} tID Selected transaction ID
 * @param {string} endpoint Endpoint to which request is made to get transaction
 * @returns None
 */

const getAndLoadTForEdit = async function (tID, endpoint) {
	const url = `${endpoint}?id=${tID}`;
	const res = await makeFetchRequest(url);

	if (!res.date) return showError('Error loading the transaction!');

	// load values to form
	dateEl.value = res.date;
	descEl.value = res.desc;
	amountEl.value = res.amount;
	paidToEl.value = res.paid_to;
	modeEl.value = res.mode;

	if (res.type === 'Income') document.getElementById('income').checked = true;
	else document.getElementById('expense').checked = true;

	// change btn text content
	submitBtn.textContent = 'Update transaction';

	// storing the updated transaction id to sessionStorage to fetch later on reload for showing changes
	sessionStorage.setItem(UPDATE_TID, res._id);
};

/**
 * Delete transaction
 *
 * @param {object} tRow Transaction row from html - table
 * @param {string} tID Selected transaction ID
 * @param {string} endpoint Endpoint to which request will be made
 */

const initiateDeleteT = function (tRow, tID, endpoint) {
	// get child elements from tRow (transaction) who have no children and strike them before deleting that transaction
	Array.from(tRow.children).forEach(field => {
		if (field.children.length > 0) return;

		// to keep the amount text color same as the original
		if (field.classList.contains('t-type-income'))
			field.innerHTML = getStrikeSpanMarkup(field.innerHTML, 't-type-income');
		else if (field.classList.contains('t-type-expense'))
			field.innerHTML = getStrikeSpanMarkup(field.innerHTML, 't-type-expense');
		else field.innerHTML = getStrikeSpanMarkup(field.innerHTML, 'strike-text');
	});

	// make ajax call to delete the transaction after a timeout
	const reqUrl = `${endpoint}?id=${tID}`;
	setTimeout(async () => {
		const res = await makeFetchRequest(reqUrl);
		if (!res.success) return showError('Error deleting transaction!');

		window.location.reload();
	}, 1000);
};

/**
 * Flash the selected transaction
 * @param {string} tID Transaction ID of transaction that will be highlighted
 */

const highlightT = function (tID) {
	const callSetTimeoutForUpdate = (tds, ms) =>
		setTimeout(() => tds.forEach(el => el.classList.toggle('updated-t')), ms);

	const tRow = document.getElementById(tID);

	// selecting only the td where there is text
	const tds = [];
	Array.from(tRow.children).forEach(field => {
		if (field.children.length <= 0) tds.push(field);
	});

	callSetTimeoutForUpdate(tds, 500);
	callSetTimeoutForUpdate(tds, 1000);
};

/**
 * Generate new markup for strikethrough text
 *
 * @param {string} currMarkup Existing markup of row element in table
 * @param {string} spanClass Class name of the span class that will be added to the new markup
 * @returns {string} New markup with strikethrough text
 */

const getStrikeSpanMarkup = (currMarkup, spanClass) =>
	`<s class="strike"><span class="${spanClass}">${currMarkup}</span></s>`;

const showModal = () => (modal.style.display = 'block');
const hideModal = () => (modal.style.display = 'none');
