'use strict';

// EL = EventListener
// El = Element

// main dom elements

// modal
const modalContainer = document.querySelector('.modal-container');
const modalCloseBtn = document.querySelector('.modal-close-btn');
const overlay = document.querySelector('.overlay');

// modal variables
const modalTitle = document.querySelector('.modal-title');
const modalDate = document.querySelector('.modal-date');
const modalTextContainer = document.querySelector('.modal-text-container');

//
//
//

// navbar
const navigationLinks = document.querySelectorAll('[data-nav-link]');
const pages = document.querySelectorAll('[data-page]');

// page names
const addPage = 'add';
const reportPage = 'report';
const settingsPage = 'settings';

//
//
//

// under ADD tab
// transaction form
const tForm = document.getElementById('t-form');
const formTitle = document.querySelector('.t-form-heading');

// t-form element ids
const [
	dateID, //               date
	descID, //               text
	amountID, //             number
	modeID, //               select dropdown
	typeInputGroupName, //   radio group
	paidToID, //             text
] = ['date', 'desc', 'amount', 'mode', 'type', 'paid_to'];

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

const submitBtn = document.querySelector('.btn-add');
const clearBtn = document.querySelector('.btn-clear');

// error on form submissions

const errDivAddPage = document.querySelector('.error-add-page');
const errDivReportPage = document.querySelector('.error-report-page');
const errDivSettingsPage = document.querySelector('.error-settings-page');

let currEditTID = '';

const table = document.querySelector('.t-table');

//
//
//

// under REPORT tab
// sort-form
const sortForm = document.getElementById('sort-form');

// sort-form element ids
const viewID = 'view';

// sort-form select option element values (view)

const [
	viewAll,
	viewLast7Days,
	viewLast30Days,
	viewThisMonth,
	viewLastMonth,
	viewCustom,
] = [
	'1_all',
	'2_last7',
	'3_last30',
	'4_this_month',
	'5_last_month',
	'6_custom',
];

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

const sortBtn = document.querySelector('.btn-sort');

//
//
//

// settings page
const currencyID = 'currency';
const modesOfPaymentID = 'modes';
const monthlyBudgetID = 'monthly_budget';

const settingsForm = document.getElementById('set-form');
const setBtn = document.querySelector('.btn-set');

const allModesOfPayment = new Map([
	// value: shown on UI
	['1_cash', 'Cash'],
	['2_phonepe', 'PhonePe'],
	['3_googlepay', 'Google Pay'],
	['4_paytm', 'PayTM'],
	['5_card', 'Card'],
	['6_other', 'Other'],
]);

//
//
//

// CONFIG variables
// endpoints

const [
	HOME_ENDPOINT,
	GET_ENDPOINT,
	ADD_ENDPOINT,
	EDIT_ENDPOINT,
	DEL_ENDPOINT,
	SETTINGS_ENDPOINT,
] = ['/', '/get', '/add', '/edit', '/del', '/settings'];

// session storage keys
// to preserve the curr page on reload
const currentPage = 'currPage';

// global class names
const cHidden = 'hidden';
const cActive = 'active';

const cTTypeIncome = 't-type-income';
const cTTypeExpense = 't-type-expense';

const cModalFieldLabel = 'modal-field-label';
const cModalFieldValue = 'modal-field-value';

const cT = 't';
const cViewIcon = 'view-icon';
const cEditIcon = 'edit-icon';
const cDeleteIcon = 'delete-icon';

// timeouts (ms)
const errShowTimeout = 2500;

// text (messages/errors)
const errInsertT = 'Error inserting the transaction!';
const errUpdateT = 'Error updating the transaction!';
const errLoadT = 'Error loading the transaction!';
const errDeleteT = 'Error deleting the transaction!';
const errSaveSettings = 'Error saving your settings!';

// form titles
const formTitleOnAddExpense = 'Add expense';
const formTitleOnUpdateExpense = 'Update transaction';

// btn texts
const btnTextAddT = 'Add transaction';
const btnTextUpdateT = 'Update transaction';
const btnTextClear = 'Clear all';
const btnTextCancel = 'Cancel';
const btnSave = 'Save';
