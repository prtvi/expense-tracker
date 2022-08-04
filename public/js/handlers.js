const hideModalOnClick = function (e) {
	// modal here is the bg that is darkened
	if (e.target === modal && modal.style.display === 'block') hideModal();
};

const hideModalOnKeydown = function (e) {
	// looking for escape key pressed to hide modal
	if (e.key === 'Escape' && modal.style.display === 'block') hideModal();
};

const tableEL = function (e) {
	// using event delegation to add event listener to the entire table rather than every transaction
	const tRow = e.target.closest(`.${cT}`);
	if (!tRow) return;

	const tID = tRow.getAttribute('id');

	// looking for edit event
	if (e.target.classList.contains(cEditIcon))
		getAndLoadTForEdit(tID, GET_ENDPOINT);

	// delete event
	if (e.target.classList.contains(cDeleteIcon))
		initiateDeleteT(tRow, tID, DEL_ENDPOINT);

	// view event
	if (e.target.classList.contains(cViewIcon)) displayTModal(tID, GET_ENDPOINT);

	// double click event
	if (e.type === 'dblclick') displayTModal(tID, GET_ENDPOINT);
};

const tFormEL = function (e) {
	e.preventDefault();

	// if btn text-content is for adding transaction then send form data wo update options
	if (submitBtn.textContent === btnTextAddT) sendFormData(tForm, ADD_ENDPOINT);
	// else send UPDATE_TID
	else {
		// to check on reload if update transaction btn was clicked or page was reloaded
		sessionStorage.setItem(UPDATE_TRUE, 'true');
		const tID = sessionStorage.getItem(UPDATE_TID);

		if (tID) sendFormData(tForm, EDIT_ENDPOINT, true, tID);
		else showError(errUpdateT);
	}
};

const updateTHandler = function (e) {
	/*
	UPDATE process
	- event listener for the table checks if edit icon is clicked
	- get the transaction id 
	- fetch the document (transaction) from the db and display for edit
	- store this document id in sessionStorage to fetch after reload
	- look for form submission, if the btn text is 'Update transaction' then send update_id and data and update the transaction
	- use the stored tID to render the update
*/
	const tID = sessionStorage.getItem(UPDATE_TID);
	const updateTrue = sessionStorage.getItem(UPDATE_TRUE);

	// render update animation only when reloading happens through update transaction btn
	if (tID && updateTrue) {
		highlightT(tID);
		sessionStorage.removeItem(UPDATE_TID);
		sessionStorage.removeItem(UPDATE_TRUE);
	}
};

const sortFormEL = function (e) {
	// submit form by adding the form params into url rather than making a request to "/" route
	// this is done to use sort middleware rather than making a get request to a route and waiting for content to arrive on that route
	// here the middleware picks the request params from the url and renders the page using only "/" route
	e.preventDefault();
	window.location.href = generateQueryUrl(sortForm, HOME_ENDPOINT);
};

const viewEL = function (e) {
	// enable or disable custom dates container based on view selection
	if (this.value === viewCustom) enableCustomDatesContainer(true);
	else enableCustomDatesContainer(false);
};

const sortParamsLoader = function (e) {
	// on sort-form submission preserve the sort option and display the same
	const sortParams = new URLSearchParams(window.location.search);

	// sort by options
	// set the sort input element
	if (sortParams.get(sortID) === sortAscID) sortEl.value = sortAscID;
	else if (sortParams.get(sortID) === sortDesID)
		sortEl.value = sortDesID;

	// set the sortInput element value
	viewEl.value = sortParams.get(viewID) || viewAll;

	// if sortParam is not custom then return
	if (sortParams.get(viewID) !== viewCustom) return;

	// else enable the custom dates container & set the corresponding values from sortParams
	enableCustomDatesContainer(true);
	customDateStartEl.value = sortParams.get(customDateStartID);
	customDateEndEl.value = sortParams.get(customDateEndID);
};

// modal close event handlers
window.addEventListener('click', hideModalOnClick);
window.addEventListener('keydown', hideModalOnKeydown);
modalClose.addEventListener('click', hideModal);

// table
table && ['click', 'dblclick'].forEach(e => table.addEventListener(e, tableEL));

// t-form
tForm.addEventListener('submit', tFormEL);
window.addEventListener('load', updateTHandler);
clearBtn.addEventListener('click', changeFormLabels);

// change paid_to text to "self" on type income, else ''
typeIncomeEl.addEventListener('input', () => (paidToEl.value = 'self'));
typeExpenseEl.addEventListener('input', () => (paidToEl.value = ''));

// sort-form
sortForm.addEventListener('submit', sortFormEL);
viewEl.addEventListener('input', viewEL);
window.addEventListener('load', sortParamsLoader);

// to prevent view_end_date < view_start_date
customDateStartEl.addEventListener('input', () =>
	customDateEndEl.setAttribute('min', customDateStartEl.value)
);
