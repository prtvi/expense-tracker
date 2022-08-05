'use strict';

// EL = EventListener
// El = Element

// main dom elements

const tForm = document.getElementById('t-form');
const table = document.querySelector('.t-table');

const errDiv = document.querySelector('.error-div');
const errText = document.querySelector('.error-text');

// t-form element ids
const dateID = 'date'; //             date
const descID = 'desc'; //             text
const amountID = 'amount'; //         number
const modeID = 'mode'; //             select dropdown
const typeInputGroupName = 'type'; // radio group
const paidToID = 'paid_to'; //        text

// t-form "type" (5) ids & values
const typeIncomeID = 'income';
const typeExpenseID = 'expense';

// t-form elements, except type
const dateEl = document.getElementById(dateID);
const descEl = document.getElementById(descID);
const amountEl = document.getElementById(amountID);
const modeEl = document.getElementById(modeID);
const paidToEl = document.getElementById(paidToID);

// t-form type options (income & expense)
const typeIncomeEl = document.getElementById(typeIncomeID);
const typeExpenseEl = document.getElementById(typeExpenseID);

//
//
//
//

// sort-form
const sortForm = document.getElementById('sort-form');

// sort-form element ids
const viewID = 'view';

// sort-form select option element values (view)
const viewAll = '1';
const viewLast7Days = '2';
const viewLast30Days = '3';
const viewThisMonth = '4';
const viewLastMonth = '5';
const viewCustom = '6';

// for when view element is chosen as "custom"
const customDateStartID = 'date_start';
const customDateEndID = 'date_end';

// type select
const sortID = 'sort';

// id and values
const sortAscID = 'asc';
const sortDesID = 'des';

// sort-form elements
const viewEl = document.getElementById(viewID);
const customDatesContainer = document.querySelector('.custom-dates-container');
const customDateStartEl = document.getElementById(customDateStartID);
const customDateEndEl = document.getElementById(customDateEndID);

// sort by asc/des option
const sortEl = document.getElementById(sortID);
const sortAscEl = document.getElementById(sortAscID);
const sortDesEl = document.getElementById(sortDesID);

// endpoints
const HOME_ENDPOINT = '/';
const GET_ENDPOINT = '/get';
const ADD_ENDPOINT = '/add';
const EDIT_ENDPOINT = '/edit';
const DEL_ENDPOINT = '/del';

// keys for session storage
const UPDATE_TID = 'UPDATE_TID';
const UPDATE_TRUE = 'UPDATE';

// global class names
const cHidden = 'hidden';

const cStrike = 'strike';
const cStrikeText = 'strike-text';

const cTTypeIncome = 't-type-income';
const cTTypeExpense = 't-type-expense';

const cUpdatedT = 'updated-t';

const cModalTitle = 'modal-title';
const cModalTFieldDiv = 'modal-t-field-div';
const cModalTFieldLabel = 'modal-t-field-label';
const cModalTFieldValue = 'modal-t-field-value';

const cT = 't';
const cViewIcon = 'view-icon';
const cEditIcon = 'edit-icon';
const cDeleteIcon = 'delete-icon';

// timeouts (ms)
const errShowTimeout = 2500;
const updateTTimeout = 500;
const updateTTimeout2 = 1000;
const deleteTTimeout = 1000;
const clearUrlTimeout = 3000;

// text (messages/errors)
const errInsertT = 'Error inserting the transaction!';
const errUpdateT = 'Error updating the transaction!';
const errLoadT = 'Error loading the transaction!';
const errDeleteT = 'Error deleting the transaction!';

// form titles
const formTitleOnAddExpense = 'Add expense';
const formTitleOnUpdateExpense = 'Update transaction';

const formTitle = document.querySelector('.t-form-heading');

// btn texts
const btnTextAddT = 'Add transaction';
const btnTextUpdateT = 'Update transaction';
const btnTextClear = 'Clear all';
const btnTextCancel = 'Cancel';

// btns
const submitBtn = document.querySelector('.btn-add');
const clearBtn = document.querySelector('.btn-clear');
const sortBtn = document.querySelector('.btn-sort');

// modal
const modal = document.querySelector('.modal');
const modalContent = document.querySelector('.modal-content');
const modalClose = document.querySelector('.close-modal-span');

// functions

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

  mode ? (modeEl.value = 'PhonePe') : (modeEl.value = 'Google Pay');

  type ? (typeIncomeEl.checked = true) : (typeExpenseEl.checked = true);
})();

/**
 * Render error poput with given "errMsg"
 * @param {string} errMsg Error message to display
 */
const showError = function (errMsg) {
  errText.textContent = errMsg;
  errDiv.classList.remove(cHidden);
  setTimeout(() => errDiv.classList.add(cHidden), errShowTimeout);
};

/**
 * On success event from form submission
 *
 * @param {boolean} reload If true, then reload page
 */
const onFormSubmitSuccess = reload => {
  reload ? window.location.reload() : console.log('debugging (no reload)');
};

/**
 * Render error on form submission failure
 * @param {string} endpoint Error based on endpoint
 */
const onFormSubmitFailure = function (endpoint) {
  if (endpoint === ADD_ENDPOINT) showError(errInsertT);
  if (endpoint === EDIT_ENDPOINT) showError(errUpdateT);
};

/**
 * Generate new markup for strikethrough text on table while deleting transaction
 *
 * @param {string} currMarkup Existing markup of row element in table
 * @param {string} spanClass Class name of the span class that will be added to the new markup
 * @returns {string} New markup with strikethrough text
 */
const getStrikeSpanMarkup = (currMarkup, spanClass) => {
  return `<s class="${cStrike}"><span class="${spanClass}">${currMarkup}</span></s>`;
};

// self explanatory
const showModal = () => (modal.style.display = 'block');
const hideModal = () => (modal.style.display = 'none');

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
  if (res.success) onFormSubmitSuccess(true);
  // else render errors accordingly
  else onFormSubmitFailure(endpoint);
};

const changeFormLabels = function (add = true) {
  // if ADD
  if (add) {
    formTitle.textContent = formTitleOnAddExpense;
    submitBtn.textContent = btnTextAddT;
    clearBtn.textContent = btnTextClear;
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

  if (!res.date) return showError(errLoadT);

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
    if (field.classList.contains(cTTypeIncome))
      field.innerHTML = getStrikeSpanMarkup(field.innerHTML, cTTypeIncome);
    else if (field.classList.contains(cTTypeExpense))
      field.innerHTML = getStrikeSpanMarkup(field.innerHTML, cTTypeExpense);
    else field.innerHTML = getStrikeSpanMarkup(field.innerHTML, cStrikeText);
  });

  // make ajax call to delete the transaction after a timeout
  const reqUrl = `${endpoint}?id=${tID}`;
  setTimeout(async () => {
    const res = await makeFetchRequest(reqUrl);
    if (!res.success) return showError(errDeleteT);

    window.location.reload();
  }, deleteTTimeout);
};

/**
 * Flash the selected transaction
 * @param {string} tID Transaction ID of transaction that will be highlighted
 */
const highlightT = function (tID) {
  // local function for highlightT function
  const callSetTimeoutForUpdate = (tds, ms) =>
    setTimeout(() => tds.forEach(el => el.classList.toggle(cUpdatedT)), ms);

  const tRow = document.getElementById(tID);

  // selecting only the td where there is text
  const tds = [];
  Array.from(tRow.children).forEach(field => {
    if (field.children.length <= 0) tds.push(field);
  });

  callSetTimeoutForUpdate(tds, updateTTimeout);
  callSetTimeoutForUpdate(tds, updateTTimeout2);
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

/**
 * Displays a modal with the data from the selected transaction, after fetching it from database
 *
 * @param {string} tID Transaction ID of selected transaction to be fetched
 * @param {string} endpoint Endpoint to which request will be made
 * @returns None
 */
const displayTModal = async function (tID, endpoint) {
  const url = `${endpoint}?id=${tID}`;
  const res = await makeFetchRequest(url);

  if (!res.date) return showError(errLoadT);

  // format date in readable format
  const formattedDate = formatDate(res.date);

  // TODO: add currency symbol here
  // TODO: truncated desc

  const typeUpper = res.type.slice(0, 1).toUpperCase() + res.type.slice(1);

  // map to create modal content
  const modalContentMap = new Map([
    ['Date', formattedDate],
    ['Description', res.desc],
    ['Amount', res.amount],
    ['Mode', res.mode],
    ['Type of transaction', typeUpper],
    ['Paid to', res.paid_to],
  ]);

  // generate dom for modal content
  let fieldContainers = '';
  modalContentMap.forEach((value, key) => {
    fieldContainers += `<div class="${cModalTFieldDiv}">
	<label class="${cModalTFieldLabel}">${key}</label>
	<p class="${cModalTFieldValue}">${value}</p>
	</div>`;
  });

  // attach title & modal content
  modalContent.innerHTML =
    `<h3 class="${cModalTitle}">${res.desc} on ${formattedDate}</h3>` +
    fieldContainers;

  // display modal
  showModal();
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
