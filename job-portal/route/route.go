package route

const(
	REGISTER_APPLICANT = "/register/applicant"
	LOGIN_APPLICANT = "/login/applicant"

	REGISTER_COMPANY = "/register/company"
	LOGIN_COMPANY = "/login/company"

	ACTIVATE_APPLICANT = "/applicant/{userId}"
	ACTIVATE_COMPANY = "/company/{userId}"

	APPLICANT_PROFILE = "/profile/applicants"
	COMPANY_PROFILE = "/profile/companies"

	JOBS = "/jobs" // GET ALL JOBS & POST JOB FOR COMPANY
	JOBS_RECOMENDATION = "/jobs/recommendation" // JOB RECOMMENDATION FOR APPLICANT
	JOB_MANIPULATION = "/jobs/{jobId}" // GET DETAIL JOB,DELETE JOB, UPDATE
	JOB_TAKEDOWN = "/jobs/{jobId}/takedown" // TEMPORARY TAKE DOWN

	APPLY = "/jobs/{jobId}/apply" // APPLY JOB

	MY_JOBS = "/jobs/applied"
	MY_JOB_MANIPULATION = "/jobs/{jobId}/applied" // CANCEL PROPOSE, UPDATE PROPOSE DATA

	COMPANY_JOBS = "/jobs/{companyId}" // GET ALL POSTED JOBS BY COMPANY
	JOBSS_APPLICANT = "/jobs/{jobId}/applicants"
	JOBS_APPLICANT_MANIPULATION = "/jobs/{jobId}/applicants/{applicanId}" // REJECT APPLICANT OR GIVE INFORMATION VIA EMAIL

)
