const { Octokit } = require("@octokit/rest");
const env = require("./.env");

const GITHUB_API_TOKEN_ISSUES = env.GITHUB_API_TOKEN_ISSUES;
// Set up Octokit instance with authentication token
const octokit = new Octokit({ auth: GITHUB_API_TOKEN_ISSUES });
const owner = "juliocnsouzadev";
const repo = "board";

// Function to create a new issue
async function createIssue(title, body) {
    const { data: issue } = await octokit.issues.create({
        owner,
        repo,
        title,
        body,
    });
    return issue.html_url;
}

// Function to read an existing issue
async function readIssue(issueNumber) {
    const { data: issue } = await octokit.issues.get({
        owner,
        repo,
        issue_number: issueNumber,
    });
    return {
        title: issue.title,
        body: issue.body,
        url: issue.html_url,
    };
}

// Function to read an existing issue
async function readAllIssues() {
    const { data: issue } = await octokit.issues.list({
        owner,
        repo,
    });
    return [
        {
            title: issue.title,
            body: issue.body,
            url: issue.html_url,
        },
    ];
}

// Function to update an existing issue
async function updateIssue(issueNumber, title, body) {
    const { data: issue } = await octokit.issues.update({
        owner,
        repo,
        issue_number: issueNumber,
        title,
        body,
    });
    return issue.html_url;
}

// Function to delete an existing issue
async function deleteIssue(issueNumber) {
    const { data: issue } = await octokit.issues.delete({
        owner,
        repo,
        issue_number: issueNumber,
    });
    return issueNumber;
}

async function main() {
    try {
        console.log("Starting issues CRUD example...");
        console.log(GITHUB_API_TOKEN_ISSUES);
        //list all issues
        const issues = await readAllIssues();
        console.log(`Issues read: ${issues}`);

        return;

        // Create a new issue
        const issueUrl = await createIssue(
            "New issue",
            "This is the issue body."
        );
        console.log(`Issue created: ${issueUrl}`);

        // Read an existing issue
        const issueNumber = /\/issues\/(\d+)/.exec(issueUrl)[1];
        const issue = await readIssue(issueNumber);
        console.log("Issue read:");
        console.log(issue);

        // Update an existing issue
        const updatedIssueUrl = await updateIssue(
            issueNumber,
            "Updated issue",
            "This is the updated issue body."
        );
        console.log(`Issue updated: ${updatedIssueUrl}`);

        // Delete an existing issue
        const deletedIssueNumber = await deleteIssue(issueNumber);
        console.log(`Issue deleted: ${deletedIssueNumber}`);
    } catch (error) {
        console.error(error);
    }
}

main();
