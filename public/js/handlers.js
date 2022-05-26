const tableEventListener = function (e) {
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

const tFormEventListener = function (e) {
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

const updateProcessHandler = function (e) {
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

tForm.addEventListener('submit', tFormEventListener);
window.addEventListener('load', updateProcessHandler);

table && table.addEventListener('click', tableEventListener);
table && table.addEventListener('dblclick', tableEventListener);

// modal close event handlers
window.addEventListener('click', e => {
	if (e.target === modal && modal.style.display === 'block') hideModal();
});

window.addEventListener('keydown', e => {
	if (e.key === 'Escape' && modal.style.display === 'block') hideModal();
});

modalClose.addEventListener('click', hideModal);