## The game.
1. Add the Male team.
1. Open registration to the Male team.
1. Add members to the Male team.
1. Close registration to the Male team.
1. Add Female team.
1. Open registration to the Female team.
1. Add members to the Female team.
1. Close registration to the Female team.
1. Know each other.
1. Vote for each other.
1. Receive matches.
1. Delete data.

## Actions.
1. Start/End a session. [admin]
1. CRUD Team. [admin]
1. Open/close registration to the team. [admin]
1. CRUD Team member [admin,user]
1. Vote. [user]
1. Receive matches. [user]

## Restrictions.
1. Only one session at a time for the admin.
1. Only two teams per session: Male and Female.
1. User can be registered only for one team.
1. User can edit his profile only when registration to the team is open.
1. User can vote only once per session. Though, multiple choices are allowed.

## Session states.
1. Teams creation.
1. Team members registration.
1. Speed dating.
1. Voting.
1. Match making. [when all players vote, or by admin]

## Corner cases.
1. Some player never votes. Then the admin can start the match making.
