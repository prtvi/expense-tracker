'use strict';

// EL = EventListener
// El = Element

// main dom elements

// modal
const modal = document.querySelector('.modal');
const modalContent = document.querySelector('.modal-content');
const modalClose = document.querySelector('.close-modal-span');

// navbar
const navigationLinks = document.querySelectorAll('[data-nav-link]');
const pages = document.querySelectorAll('[data-page]');
const pageNames = ['add', 'report', 'settings'];

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

// error on form submission
const errDiv = document.querySelector('.error-div');
const errText = document.querySelector('.error-text');

let currEditTID = '';

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
] = ['1', '2', '3', '4', '5', '6'];

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

const table = document.querySelector('.t-table');

//
//
//

// CONFIG variables
// endpoints

const [HOME_ENDPOINT, GET_ENDPOINT, ADD_ENDPOINT, EDIT_ENDPOINT, DEL_ENDPOINT] =
	['/', '/get', '/add', '/edit', '/del'];

// session storage keys
// to preserve the curr page on reload
const currentPage = 'currPage';

// global class names
const cHidden = 'hidden';
const cActive = 'active';

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

// btn texts
const btnTextAddT = 'Add transaction';
const btnTextUpdateT = 'Update transaction';
const btnTextClear = 'Clear all';
const btnTextCancel = 'Cancel';
const btnSave = 'Save';
