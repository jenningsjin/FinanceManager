$(document).ready(function(){
	$("#login-form").on("submit", function(e) {
		e.preventDefault();
		$.ajax({
			type: "POST",
			url: "/login/",
			data: $(this).serialize(),
			success: function(data) {
				console.log("success");
				$(location).attr('href', '/user/')
			},
			error: function(data) {
				console.log("error");
				$("#errorMessageLogin").text(data.responseText);
			}
		});
	});

	$("#signup-form").on("submit", function(e) {
		e.preventDefault();
		$.ajax({
			type: "POST",
			url: "/signup/",
			data: $(this).serialize(),
			success: function(data) {
				console.log("success");
				$("#errorMessageSignup").text("Created User Successfully. Please Log in");
			},
			error: function(data) {
				console.log("error");
				$("#errorMessageSignup").text(data.responseText);
			}
		});
	});
});