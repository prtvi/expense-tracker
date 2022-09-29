'use strict';

//To keep the default date to "today"
(function (dateEl) {
	const newDate = new Date();
	const year = newDate.getFullYear();
	const month = String(newDate.getMonth() + 1).padStart(2, '0');
	const day = String(newDate.getDate()).padStart(2, '0');

	dateEl.value = `${year}-${month}-${day}`;
})(dateEl);

(function () {
	const desc = (Math.random() + 1).toString(36).substring(2);
	const amount = Math.floor(Math.random() * (550 - 10) + 10);
	const paidTo = (Math.random() + 1).toString(36).substring(2);
	const mode = Math.random() > 0.5 ? true : false;
	const type = Math.random() > 0.5 ? true : false;

	descEl.value = desc;
	amountEl.value = amount;
	paidToEl.value = paidTo;

	mode ? (modeEl.value = '1_cash') : (modeEl.value = '2_phonepe');

	type ? (typeIncomeEl.checked = true) : (typeExpenseEl.checked = true);
})();

const toggleModal = function () {
	overlay.classList.toggle('active');
	modalContainer.classList.toggle('active');
};

/**
 * Render error poput with given "errMsg"
 * @param {string} errMsg Error message to display
 */
const showError = function (errDiv, errMsg) {
	const errText = errDiv.querySelector('.error-text');

	errText.textContent = errMsg;

	errDiv.classList.remove(cHidden);
	setTimeout(() => errDiv.classList.add(cHidden), errShowTimeout);
};

/**
 * Render error on form submission failure
 * @param {string} endpoint Error based on endpoint
 */
const onFormSubmitFailure = function (endpoint) {
	if (endpoint === ADD_ENDPOINT) showError(errDivAddPage, errInsertT);
	if (endpoint === EDIT_ENDPOINT) showError(errDivAddPage, errUpdateT);
	if (endpoint === SETTINGS_ENDPOINT)
		showError(errDivSettingsPage, errSaveSettings);
};

/**
 * Make fetch request and get back JSON response
 *
 * @param {string} url URL / endpoint to which request will be made
 * @returns {JSON} JSON response from server
 */
const makeFetchRequest = async function (url) {
	const res = await fetch(url);
	const resJson = await res.json();

	if (res.ok) return resJson;
	return false;
};

/**
 * Function to generate request URL from form data
 *
 * @param {object} form Form element from which request URL is to be generated
 * @param {string} endpoint Endpoint to which the params will be appended
 * @returns {string} The generated request URL
 */
const generateQueryUrl = function (form, endpoint) {
	const formData = new FormData(form);
	let reqUrl = `${endpoint}?`;

	// generate query with form data
	Array.from(formData.keys()).forEach(key => {
		if (!formData.get(key)) return;
		reqUrl += `${key}=${formData.get(key)}&`;
	});

	// remove the last '&' from query string
	return reqUrl.slice(0, -1);
};

/**
 * Send current form data to backend
 * if update === false, then sends form data for insert transaction operation
 * if update === true, then sends tID & current form data to update the edited transaction
 *
 * @param {object} form Entire html form element
 * @param {string} endpoint Endpoint to which request will be made
 * @param {boolean} update Check if update or edit request is to be made
 * @param {string} tID Transaction ID of the transaction to be updated
 */
const sendFormData = async function (
	form,
	endpoint,
	update = false,
	tID = null
) {
	let reqUrl = generateQueryUrl(form, endpoint);

	// if update true then append id to query
	if (update) reqUrl += `&id=${tID}`;

	const res = await makeFetchRequest(reqUrl);

	// if success then reload window to reload transactions
	if (res.success) window.location.reload();
	// else render errors accordingly
	else onFormSubmitFailure(endpoint);
};

/**
 * Change form labels as per form submission
 * @param {boolean} add Change form titles depending on page
 */
const changeFormLabels = function (add = true) {
	// if ADD
	if (add) {
		formTitle.textContent = formTitleOnAddExpense;
		submitBtn.textContent = btnTextAddT;
		clearBtn.textContent = btnTextClear;
		switchPage(reportPage);
	} else {
		// if UPDATE
		formTitle.textContent = formTitleOnUpdateExpense;
		submitBtn.textContent = btnTextUpdateT;
		clearBtn.textContent = btnTextCancel;
	}
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

	if (!res.date) return showError(errDivAddPage, errLoadT);

	// load values to form
	dateEl.value = res.date;
	descEl.value = res.desc;
	amountEl.value = res.amount;
	modeEl.value = res.mode;
	paidToEl.value = res.paid_to;

	if (res.type === typeIncomeID) typeIncomeEl.checked = true;
	else typeExpenseEl.checked = true;

	// change btn text content
	changeFormLabels(false);

	// save id for currently editing t
	currEditTID = res._id;
};

/**
 *
 * @param {String} desc the input text
 * @param {Number} maxLen maximum length of desc returned
 * @returns the truncated string
 */
const truncateDesc = function (desc, maxLen) {
	return desc.length >= maxLen ? desc.slice(0, maxLen) + '...' : desc;
};

/**
 * Delete transaction
 *
 * @param {object} tRow Transaction row from html - table
 * @param {string} tID Selected transaction ID
 * @param {string} endpoint Endpoint to which request will be made
 */
const initiateDeleteT = function (tRow, tID, endpoint) {
	const maxDescLen = 15;

	// first confirm deletion
	const deleteModalMarkup = `<div class="del-btn-container">
	<button class="btn btn-del-modal yes">Yes</button>
	<button class="btn btn-del-modal cancel">Cancel</button>
	</div>`;

	loadModalData(
		`Delete ${truncateDesc(tRow.children[1].textContent, maxDescLen)}?`,
		'',
		deleteModalMarkup
	);

	toggleModal();

	const btnYes = document.querySelector('.yes');
	const btnCancel = document.querySelector('.cancel');

	btnCancel.addEventListener('click', toggleModal);

	btnYes.addEventListener('click', async function () {
		// make ajax call to delete the transaction
		const reqUrl = `${endpoint}?id=${tID}`;
		const res = await makeFetchRequest(reqUrl);

		if (!res.success) return showError(errDivReportPage, errDeleteT);

		window.location.reload();
	});
};

/**
 * Format date: 2022-05-25 => Wed, 25 May 2022
 *
 * @param {string} dateStr Date string of format "2022-05-25"
 * @returns {string} Date string of format "Wed, 25 May 2022"
 */
const formatDate = function (dateStr) {
	const [weekday, month, date, year] = String(new Date(dateStr))
		.slice(0, 16)
		.split(' ');

	return `${weekday}, ${date} ${month} ${year}`;
};

const getCurrencySymbol = function () {
	const symbolText = document.querySelector(
		`[data-currency-symbol]`
	).textContent;
	const [start, end] = [symbolText.indexOf('('), symbolText.indexOf(')')];

	return symbolText.slice(start + 1, end);
};

const markupForExpenseOrIncome = function (elementVal, typeT) {
	return `<span class="${
		typeT === typeExpenseID ? cTTypeExpense : cTTypeIncome
	}">${elementVal}</span>`;
};

const generateTModalMarkup = function (res) {
	const markupForAmount = markupForExpenseOrIncome(
		`${getCurrencySymbol()} ${res.amount}`,
		res.type
	);
	const markupForType = markupForExpenseOrIncome(res.type, res.type);

	// map to create modal content
	const modalContentMap = new Map([
		['Description', res.desc],
		['Amount', markupForAmount],
		['Mode', allModesOfPayment.get(res.mode)],
		['Type of transaction', markupForType],
		['Paid to', res.paid_to],
	]);

	// generate dom for modal content
	let fieldContainers = '';
	modalContentMap.forEach((value, key) => {
		fieldContainers += `<div class="modal-field"><label class="${cModalFieldLabel}">${key}</label><p class="${cModalFieldValue}">${value}</p></div>`;
	});

	return fieldContainers;
};

/**
 * Displays a modal with the data from the selected transaction, after fetching it from database
 *
 * @param {string} tID Transaction ID of selected transaction to be fetched
 * @param {string} endpoint Endpoint to which request will be made
 * @returns None
 */
const displayTModal = async function (tID, endpoint) {
	const maxModalTitleLength = 20;

	const url = `${endpoint}?id=${tID}`;
	const res = await makeFetchRequest(url);

	if (!res.date) return showError(errDivReportPage, errLoadT);

	// attach title & modal content
	const title = truncateDesc(res.desc, maxModalTitleLength);
	const date = formatDate(res.date);
	const textContent = generateTModalMarkup(res);

	loadModalData(title, date, textContent);
	toggleModal();
};

/**
 * If enable is true, then unhides the custom dates container and adds required attribute to start & end date inputs, if false, then reverses the same operation
 *
 * @param {boolean} enable True if custom dates container is to be revealed
 */
const enableCustomDatesContainer = function (enable) {
	if (enable) {
		customDatesContainer.classList.remove(cHidden);

		customDateStartEl.setAttribute('required', 'true');
		customDateEndEl.setAttribute('required', 'true');
	} else {
		customDateStartEl.removeAttribute('required');
		customDateEndEl.removeAttribute('required');

		customDatesContainer.classList.add(cHidden);
	}
};

/**
 * To switch to the given page
 * @param {String} page Switch to this page
 */
const switchPage = function (page) {
	navigationLinks.forEach(nl => nl.classList.remove(cActive));

	pages.forEach(pg => pg.classList.remove(cActive));

	navigationLinks.forEach(nl => {
		if (nl.dataset.navLink === page) nl.classList.add(cActive);
	});

	pages.forEach(pg => {
		if (pg.dataset.page === page) pg.classList.add(cActive);
	});
};

/**
 *
 * @param {String} title Title for the modal
 * @param {String} date Date for the modal
 * @param {String | HTML} textContent HTML/String content for modal
 */
const loadModalData = function (title, date, textContent) {
	modalDate.textContent = date;
	modalTitle.textContent = title;
	modalTextContainer.innerHTML = textContent;
};

const enableNavLinks = function (enable = true) {
	if (enable) navigationLinks.forEach(nl => (nl.disabled = false));
	else navigationLinks.forEach(nl => (nl.disabled = true));
};

const getCColor = function (value) {
	return value >= 0 ? cTTypeIncome : cTTypeExpense;
};

const generateSummaryModalDates = function (dateString) {
	const startDate = formatDate(
		dateString.slice(dateString.indexOf('start=') + 1, dateString.indexOf('&'))
	);
	const endDate = formatDate(dateString.slice(dateString.indexOf('end=') + 1));

	return `${startDate} to ${endDate}`;
};

const generateSummaryModalMarkup = function (res, includeMainSummary = false) {
	let tRows = ``;
	for (const [key, value] of allModesOfPayment) {
		const data = res.indi_mode_sums[key];

		if (data.income || data.expense) {
			const balance = data.income - data.expense;
			tRows += `
			<tr>
			<td>${value}</td>
			<td>${data.income}</td>
			<td>${data.expense}</td>
			<td class="last-col ${getCColor(balance)}">${balance}</td>
			</tr>`;
		}
	}

	let markup = `
	<div class="modal-table">

	${
		includeMainSummary
			? `<table class="modal-main-summary">
				<tr>
					<th>Total income</th>
					<th>Total expense</th>
					<th>Total balance</th>
				</tr>
				<tr>
					<td class="${cTTypeIncome}">${res.total_income}</td>
					<td class="${cTTypeExpense}">${res.total_expense}</td>
					<td class="${getCColor(res.total_balance)}">${res.total_balance}</td>
				</tr>
			</table>`
			: ``
	}
		<table class="modal-split-summary">
			<tr>
				<th>Mode</th>
				<th>Income</th>
				<th>Expense</th>
				<th class="last-col">Balance</th>
			</tr>
			${tRows}
			<tr class="last-row">
				<td>Total</td>
				<td>${res.total_income}</td>
				<td>${res.total_expense}</td>
				<td class="last-col ${getCColor(res.total_balance)}">${res.total_balance}</td>
			</tr>
		</table>
	</div>`;

	return markup;
};
