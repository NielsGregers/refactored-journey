<#---
title: Exchange Test
tag: exchange
api: post
connection: exchange
---
#>

write-host "Hello Exchange"


<#
Will return the first 5 mailboxes and format it as a table
#>

Get-Mailbox -ResultSize 5 | Format-Table
