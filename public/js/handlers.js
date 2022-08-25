'use strict';

// EL = EventListener
// El = Element

// EVENT LISTENERS

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
	if (e.target.classList.contains(cEditIcon)) {
		getAndLoadTForEdit(tID, GET_ENDPOINT);
		switchPage(pageNames[0]);
	}

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
	if (submitBtn.textContent.trim() === btnTextAddT)
		sendFormData(tForm, ADD_ENDPOINT);
	// else send form data with currEditTID
	else if (submitBtn.textContent.trim() === btnTextUpdateT) {
		sendFormData(tForm, EDIT_ENDPOINT, true, currEditTID);

		// if (currEditTID)
		// else showError(errUpdateT);
	}
};

const sortFormEL = function (e) {
	/*
	- submit form by adding the form params into url rather than making a request to "/" route
	- this is done to use sort middleware rather than making a get request to a route and waiting for content to arrive on that route
	- here the middleware picks the request params from the url and renders the page using only "/" route
	*/
	e.preventDefault();
	window.location.href = generateQueryUrl(sortForm, HOME_ENDPOINT);
};

// view element in sort form
const viewEL = function (e) {
	// enable or disable custom dates container based on view selection
	if (this.value === viewCustom) enableCustomDatesContainer(true);
	else enableCustomDatesContainer(false);
};

const sortParamsAndPageLoader = function (e) {
	const currPage = sessionStorage.getItem(currentPage);
	switchPage(currPage);

	// on sort-form submission preserve the sort option and display the same
	const sortParams = new URLSearchParams(window.location.search);

	// sort options
	// set the sort input element
	if (sortParams.get(sortID) === sortAscID) sortEl.value = sortAscID;
	else if (sortParams.get(sortID) === sortDesID) sortEl.value = sortDesID;

	// set the view element value
	viewEl.value = sortParams.get(viewID) || viewAll;

	// if sortParam is not custom then return
	if (sortParams.get(viewID) !== viewCustom) return;

	// else enable the custom dates container & set the corresponding values from sortParams
	enableCustomDatesContainer(true);
	customDateStartEl.value = sortParams.get(customDateStartID);
	customDateEndEl.value = sortParams.get(customDateEndID);
};

const navbarEL = function () {
	// remove active class from all links first
	navigationLinks.forEach(nl => nl.classList.remove(cActive));

	pages.forEach(page => {
		if (this.dataset.navLink === page.dataset.page) {
			// if matched then add active class to nav link
			this.classList.add(cActive);

			// display page with active class
			page.classList.add(cActive);

			window.scrollTo(0, 0);
		} else page.classList.remove(cActive);
	});

	// use session storage keys to preserve the currPage
	sessionStorage.setItem(currentPage, this.dataset.navLink.trim());
};

// modal close event handlers
window.addEventListener('click', hideModalOnClick);
window.addEventListener('keydown', hideModalOnKeydown);
modalClose.addEventListener('click', hideModal);

// table
table && ['click', 'dblclick'].forEach(e => table.addEventListener(e, tableEL));

// t-form
tForm.addEventListener('submit', tFormEL);
clearBtn.addEventListener('click', changeFormLabels);

// change paid_to text to "self" on type income, else ''
typeIncomeEl.addEventListener('input', () => (paidToEl.value = 'self'));
typeExpenseEl.addEventListener('input', () => (paidToEl.value = ''));

// sort-form
sortForm.addEventListener('submit', sortFormEL);
viewEl.addEventListener('input', viewEL);
window.addEventListener('load', sortParamsAndPageLoader);

// to prevent view_end_date < view_start_date
customDateStartEl.addEventListener('input', () =>
	customDateEndEl.setAttribute('min', customDateStartEl.value)
);

// to switch pages
navigationLinks.forEach(navLink => navLink.addEventListener('click', navbarEL));
