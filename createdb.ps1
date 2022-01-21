
$assemblylist =   
"Microsoft.SqlServer.Management.Common",  
"Microsoft.SqlServer.Smo",  
"Microsoft.SqlServer.Dmf ",  
"Microsoft.SqlServer.Instapi ",  
"Microsoft.SqlServer.SqlWmiManagement ",  
"Microsoft.SqlServer.ConnectionInfo ",  
"Microsoft.SqlServer.SmoExtended ",  
"Microsoft.SqlServer.SqlTDiagM ",  
"Microsoft.SqlServer.SString ",  
"Microsoft.SqlServer.Management.RegisteredServers ",  
"Microsoft.SqlServer.Management.Sdk.Sfc ",  
"Microsoft.SqlServer.SqlEnum ",  
"Microsoft.SqlServer.RegSvrEnum ",  
"Microsoft.SqlServer.WmiEnum ",  
"Microsoft.SqlServer.ServiceBrokerEnum ",  
"Microsoft.SqlServer.ConnectionInfoExtended ",  
"Microsoft.SqlServer.Management.Collector ",  
"Microsoft.SqlServer.Management.CollectorEnum",  
"Microsoft.SqlServer.Management.Dac",  
"Microsoft.SqlServer.Management.DacEnum",  
"Microsoft.SqlServer.Management.Utility",
"Microsoft.SqlServer.Management.Smo"
foreach ($asm in $assemblylist) {  
    $asm = [Reflection.Assembly]::LoadWithPartialName($asm)  
} 
$dbname = "vconnect"
$dbInstance = $env:COMPUTERNAME;
# Create a SQL Server database object
$conn = new-object Microsoft.SqlServer.Management.Common.ServerConnection($dbInstance, "sa", "Kottai2050$")
$srv = New-Object Microsoft.SqlServer.Management.Smo.Server($conn)

    $db = New-Object Microsoft.SqlServer.Management.Smo.Database($srv, $dbname)

    # Create the database
    $db.Create()




$dbname = "billing"
$dbInstance = $env:COMPUTERNAME;
# Create a SQL Server database object
$conn = new-object Microsoft.SqlServer.Management.Common.ServerConnection($dbInstance, "sa", "Kottai2050$")
$srv = New-Object Microsoft.SqlServer.Management.Smo.Server($conn)

    $db = New-Object Microsoft.SqlServer.Management.Smo.Database($srv, $dbname)

    # Create the database
    $db.Create()
$dbname = "hybr"
$dbInstance = $env:COMPUTERNAME;
# Create a SQL Server database object
$conn = new-object Microsoft.SqlServer.Management.Common.ServerConnection($dbInstance, "sa", "Kottai2050$")
$srv = New-Object Microsoft.SqlServer.Management.Smo.Server($conn)

    $db = New-Object Microsoft.SqlServer.Management.Smo.Database($srv, $dbname)

    # Create the database
    $db.Create()
