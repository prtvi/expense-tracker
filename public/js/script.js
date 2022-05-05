// to keep the default date to "today"
const formatDate = function () {
	const date = document.getElementById('date');

	const newDate = new Date();
	const year = newDate.getFullYear();
	const month = String(newDate.getMonth() + 1).padStart(2, '0');
	const day = String(newDate.getDate()).padStart(2, '0');

	date.value = `${year}-${month}-${day}`;
};

formatDate();
