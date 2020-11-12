function getAll(entity) {
	fetch('api/' + entity)
		.then((data) => {
			fetch('/template/list/' + entity + '.html')
				.then((response) => response.text())
				.then((template) => {
					var rendered = Mustache.render(template, data);
					document.getElementById('content').innerHTML = rendered;
				});
		})
}

function getById(query, entity) {
	var params = new URLSearchParams(query);
	fetch('api/' + entity + '/' + params.get('id'))
		.then((data) => {
			fetch('/template/detail/' + entity + '.html')
				.then((response) => response.text())
				.then((template) => {
					var rendered = Mustache.render(template, data);
					document.getElementById('content').innerHTML = rendered;
				});
		})
}

function home() {
	fetch('/template/home.html')
		.then((response) => response.text())
		.then((template) => {
			var rendered = Mustache.render(template, {});
			document.getElementById('content').innerHTML = rendered;
		});
}

function init() {
	router = new Navigo(null, false, '#!');
	router.on({
		'/books': function() {
			getAll('books');
		},
		'/authors': function() {
			getAll('authors');
		},
		'/pushishers': function() {
			getAll('publishers');
		},
		'/bookById': function(_, query) {
			getById(query, 'books');
		},
		'/authorById': function(_, query) {
			getById(query, 'authors');
		},
		'/publisherById': function(_, query) {
			getById(query, 'publishers');
		}
	});
	router.on(() => home());
	router.resolve();
}
