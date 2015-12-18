# Pathfinder API

# Group Paths

## generate paths [/paths{?from,to}]

### GET
Generates random paths between two locations.

+ Parameters
	+ from (required, string)- starting location
	+ to (required, string) - final location

+ Response 200 (application/json)

		{
		  "paths": [
		    {
		      "edges": [
		        {
		          "voyage": "0400S",
		          "origin": "SESTO",
		          "destination": "DEHAM",
		          "departure": "2015-11-14T23:19:07.249273439+01:00",
		          "arrival": "2015-11-15T18:19:07.249273439+01:00"
		        },
		        {
		          "voyage": "0100S",
		          "origin": "DEHAM",
		          "destination": "FIHEL",
		          "departure": "2015-11-17T17:57:07.249273439+01:00",
		          "arrival": "2015-11-18T21:59:07.249273439+01:00"
		        }
		      ]
		    },
		    {
		      "edges": [
		        {
		          "voyage": "0200T",
		          "origin": "SESTO",
		          "destination": "SEGOT",
		          "departure": "2015-11-14T16:44:07.249306459+01:00",
		          "arrival": "2015-11-15T11:35:07.249306459+01:00"
		        },
		        {
		          "voyage": "0200T",
		          "origin": "SEGOT",
		          "destination": "FIHEL",
		          "departure": "2015-11-17T19:05:07.249306459+01:00",
		          "arrival": "2015-11-18T22:54:07.249306459+01:00"
		        }
		      ]
		    },
		    {
		      "edges": [
		        {
		          "voyage": "0100S",
		          "origin": "SESTO",
		          "destination": "DEHAM",
		          "departure": "2015-11-14T12:22:07.249318801+01:00",
		          "arrival": "2015-11-15T15:51:07.249318801+01:00"
		        },
		        {
		          "voyage": "0301S",
		          "origin": "DEHAM",
		          "destination": "FIHEL",
		          "departure": "2015-11-17T15:06:07.249318801+01:00",
		          "arrival": "2015-11-18T07:38:07.249318801+01:00"
		        }
		      ]
		    },
		    {
		      "edges": [
		        {
		          "voyage": "0200T",
		          "origin": "SESTO",
		          "destination": "USNYC",
		          "departure": "2015-11-14T19:08:07.249330441+01:00",
		          "arrival": "2015-11-15T17:25:07.249330441+01:00"
		        },
		        {
		          "voyage": "0100S",
		          "origin": "USNYC",
		          "destination": "AUMEL",
		          "departure": "2015-11-17T11:13:07.249330441+01:00",
		          "arrival": "2015-11-18T17:30:07.249330441+01:00"
		        },
		        {
		          "voyage": "0300A",
		          "origin": "AUMEL",
		          "destination": "JNTKO",
		          "departure": "2015-11-20T21:11:07.249330441+01:00",
		          "arrival": "2015-11-22T02:08:07.249330441+01:00"
		        },
		        {
		          "voyage": "0301S",
		          "origin": "JNTKO",
		          "destination": "CNHGH",
		          "departure": "2015-11-24T05:52:07.249330441+01:00",
		          "arrival": "2015-11-25T12:56:07.249330441+01:00"
		        },
		        {
		          "voyage": "0400S",
		          "origin": "CNHGH",
		          "destination": "DEHAM",
		          "departure": "2015-11-27T23:49:07.249330441+01:00",
		          "arrival": "2015-11-29T01:51:07.249330441+01:00"
		        },
		        {
		          "voyage": "0400S",
		          "origin": "DEHAM",
		          "destination": "FIHEL",
		          "departure": "2015-12-01T09:52:07.249330441+01:00",
		          "arrival": "2015-12-02T07:09:07.249330441+01:00"
		        }
		      ]
		    }
		  ]
		}
