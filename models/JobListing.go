package models

type Joblisting {
	Jobtitle       string `json:"job_title"`
	Jobdescription string `json:"job_description"`
	Language       string `json:"language"`
	Location       string `json:"location"`
	Payment        string `json:"payment"`
	Url            string `json:"payment"`
}

func CreateJobListing(jobtitle, jobdescription, language, location, payment, url string) *Joblisting {
	return &Joblisting{
		Jobtitle:       jobtitle,
		Jobdescription: jobdescription,
		Language:       language,
		Location:       location,
		Payment:        payment,
		Url:            url,
	}
}
