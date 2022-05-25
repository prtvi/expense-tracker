// all global variables declared in this file

const form = document.getElementById('form');
const errDiv = document.querySelector('.error-div');
const errText = document.querySelector('.error-text');
const table = document.querySelector('.t-table');

// form elements, except type
const dateEl = document.getElementById('date');
const descEl = document.getElementById('desc');
const amountEl = document.getElementById('amount');
const paidToEl = document.getElementById('paid_to');
const modeEl = document.getElementById('mode');

// btns
const submitBtn = document.querySelector('.btn-add');

// modal
const modal = document.querySelector('.modal');
const modalContent = document.querySelector('.modal-content');
const modalClose = document.querySelector('.close-modal');

// endpoints
const GET_ENDPOINT = '/get';
const ADD_ENDPOINT = '/add';
const EDIT_ENDPOINT = '/edit';
const DEL_ENDPOINT = '/del';

// keys for session storage
const UPDATE_TID = 'UPDATE_TID';
const UPDATE_TRUE = 'UPDATE';
