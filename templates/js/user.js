$(document).ready(function(){
	$("#add-transaction	").on("submit", function(e) {
		e.preventDefault();
		$.ajax({
			type: "POST",
			url: "/add-transaction/",
			data: $(this).serialize(),
			success: function(data) {
				console.log("success");
				$(location).attr('href', '/user/')
			},
			error: function(data) {
				console.log("error");
				$("#errorMessage").text(data.responseText);
			}
		});
	});

	$("#logout").click(function() {
		$.ajax({
			type: "POST",
			url: "/logout/",
			success: function(data) {
				console.log("success");
				$(location).attr('href', '/')
			},
			error: function(data) {
				console.log("error")
				$(location).attr('href', '/')
			}
		});
	});

});




