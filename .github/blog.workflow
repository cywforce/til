workflow "post draft of blog" {
  resolves = ["Slack notification"]
  on = "push"
}

action "filter PR merged" {
  uses = "actions/bin/filter@3c0b4f0e63ea54ea5df2914b4fabf383368cd0da"
  args = "merged true"
}

action "Slack notification" {
  uses = "Ilshidur/action-slack@master"
  secrets = ["SLACK_WEBHOOK"]
  args = "A new commit has been pushed."
  needs = ["filter PR merged"]
}