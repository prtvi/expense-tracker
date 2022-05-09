// send current form data to backend
// if update === false, then sends form data for insert/add operation
// if update === true, then sends UPDATE_TID to update the loaded transaction
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

	// storing the updated transaction id to sessionStorage to fetch later on reload for showing changes
	sessionStorage.setItem(UPDATE_TID, res._id);
};

const initiateDeleteT = async function (tID, endpoint) {
	const reqUrl = `${endpoint}?id=${tID}`;
	const res = await makeFetchRequest(reqUrl);

	if (!res.success) return showError('Error deleting transaction!');

	window.location.reload();
};

// using event delegation to add event listener to the entire table rather than every transaction
const tableEventListener = function (e) {
	const tRow = e.target.closest('.t');
	const tID = tRow.getAttribute('id');

	// looking for edit event
	if (e.target.classList.contains('edit-icon'))
		getAndLoadTForEdit(tID, GET_ENDPOINT);

	// delete event
	if (e.target.classList.contains('delete-icon'))
		initiateDeleteT(tID, DEL_ENDPOINT);
};

const formEventListener = function (e) {
	e.preventDefault();

	// if btn text-content is for adding transaction then send form data wo update options
	if (submitBtn.textContent === 'Add transaction')
		sendFormData(form, ADD_ENDPOINT);
	// else send UPDATE_TID
	else {
		// to check on reload if update transaction btn was clicked or page was reloaded
		sessionStorage.setItem(UPDATE_TRUE, 'true');
		const tID = sessionStorage.getItem(UPDATE_TID);

		if (tID) sendFormData(form, EDIT_ENDPOINT, true, tID);
		else showError('Error updating the transaction');
	}
};

const updateProcessHandler = function (e) {
	const tID = sessionStorage.getItem(UPDATE_TID);
	const updateTrue = sessionStorage.getItem(UPDATE_TRUE);
	if (tID && updateTrue) {
		highLightT(tID);
		sessionStorage.removeItem(UPDATE_TID);
		sessionStorage.removeItem(UPDATE_TRUE);
	}
};

table && table.addEventListener('click', tableEventListener);
form.addEventListener('submit', formEventListener);
window.addEventListener('load', updateProcessHandler);
